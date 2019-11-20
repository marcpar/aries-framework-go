// +build js,wasm

/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package jsindexeddb

import (
	"errors"
	"fmt"
	"syscall/js"
	"time"

	"github.com/hyperledger/aries-framework-go/pkg/storage"
)

var dbVersion = 1 //nolint:gochecknoglobals

// Provider jsindexeddb implementation of storage.Provider interface
type Provider struct {
}

// NewProvider instantiates Provider
func NewProvider() (*Provider, error) {
	// TODO Add unit test for IndexedDB https://github.com/hyperledger/aries-framework-go/issues/834
	return &Provider{}, nil
}

// Close closes all stores created under this store provider
func (p *Provider) Close() error {
	return nil
}

// OpenStore open store
func (p *Provider) OpenStore(name string) (storage.Store, error) {
	db, err := openDB(name)
	if err != nil {
		return nil, err
	}

	return &store{name: name, db: db}, nil
}

func openDB(name string) (*js.Value, error) {
	dbName := "aries-" + name
	req := js.Global().Get("indexedDB").Call("open", dbName, dbVersion)
	req.Set("onupgradeneeded", js.FuncOf(func(this js.Value, inputs []js.Value) interface{} {
		fmt.Printf("indexedDB create object store %s\n", name)
		m := make(map[string]interface{})
		m["keyPath"] = "key"
		this.Get("result").Call("createObjectStore", name, m)
		return nil
	}))

	v, err := getResult(req)
	if err != nil {
		return nil, fmt.Errorf("failed to open indexedDB: %w", err)
	}

	return v, nil
}

// CloseStore closes level db store of given name
func (p *Provider) CloseStore(name string) error {
	return nil
}

type store struct {
	name string
	db   *js.Value
}

// Put stores the key and the record
func (s *store) Put(k string, v []byte) error {
	if k == "" || v == nil {
		return errors.New("key and value are mandatory")
	}

	m := make(map[string]interface{})
	m["key"] = k
	m["value"] = string(v)

	req := s.db.Call("transaction", s.name, "readwrite").Call("objectStore", s.name).Call("add", m)

	_, err := getResult(req)
	if err != nil {
		return fmt.Errorf("failed to store data: %w", err)
	}

	return nil
}

// Get fetches the record based on key
func (s *store) Get(k string) ([]byte, error) {
	if k == "" {
		return nil, errors.New("key is mandatory")
	}

	req := s.db.Call("transaction", s.name).Call("objectStore", s.name).Call("get", k)

	data, err := getResult(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get data: %w", err)
	}

	if !data.Truthy() {
		return nil, storage.ErrDataNotFound
	}

	return []byte(data.Get("value").String()), nil
}

// Iterator returns iterator for the latest snapshot of the underlying db.
func (s *store) Iterator(start, limit string) storage.StoreIterator {
	// TODO Add Iterator for IndexedDB https://github.com/hyperledger/aries-framework-go/issues/833
	return nil
}

func getResult(req js.Value) (*js.Value, error) {
	onsuccess := make(chan js.Value)
	onerror := make(chan js.Value)

	req.Set("onsuccess", js.FuncOf(func(this js.Value, inputs []js.Value) interface{} {
		onsuccess <- this.Get("result")
		return nil
	}))
	req.Set("onerror", js.FuncOf(func(this js.Value, inputs []js.Value) interface{} {
		onerror <- this.Get("error")
		return nil
	}))
	select {
	case value := <-onsuccess:
		return &value, nil
	case value := <-onerror:
		return nil, fmt.Errorf("%s %s", value.Get("name").String(),
			value.Get("message").String())
	case <-time.After(3 * time.Second):
		return nil, errors.New("timeout waiting for eve")
	}
}