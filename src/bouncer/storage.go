/*
 * storage.go - Implements the storage interface (note largely lifted from teststorage)
 *
 * Bouncer is (c) 2014 Sourdough Labs Research and Development Corp.
 *
 * License: MIT (See LICENSE for details)
 */

package main

import (
	"errors"
	//	"fmt"
	"github.com/RangelReale/osin"
)

type InMemoryStorage struct {
	clients   map[string]osin.Client
	authorize map[string]*osin.AuthorizeData
	access    map[string]*osin.AccessData
	refresh   map[string]string
}

func NewInMemoryStorage() *InMemoryStorage {
	r := &InMemoryStorage{
		clients:   make(map[string]osin.Client),
		authorize: make(map[string]*osin.AuthorizeData),
		access:    make(map[string]*osin.AccessData),
		refresh:   make(map[string]string),
	}

	return r
}

func (s *InMemoryStorage) Clone() osin.Storage {
	return s
}

func (s *InMemoryStorage) Close() {
}

func (s *InMemoryStorage) GetClient(id string) (osin.Client, error) {
	//	fmt.Printf("GetClient: %s\n", id)
	if c, ok := s.clients[id]; ok {
		return c, nil
	}
	return nil, errors.New("Client not found")
}

func (s *InMemoryStorage) SetClient(id string, client osin.Client) error {
	//	fmt.Printf("SetClient: %s\n", id)
	s.clients[id] = client
	return nil
}

func (s *InMemoryStorage) SaveAuthorize(data *osin.AuthorizeData) error {
	//	fmt.Printf("SaveAuthorize: %s\n", data.Code)
	s.authorize[data.Code] = data
	return nil
}

func (s *InMemoryStorage) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
	//	fmt.Printf("LoadAuthorize: %s\n", code)
	if d, ok := s.authorize[code]; ok {
		return d, nil
	}
	return nil, errors.New("Authorize not found")
}

func (s *InMemoryStorage) RemoveAuthorize(code string) error {
	//	fmt.Printf("RemoveAuthorize: %s\n", code)
	delete(s.authorize, code)
	return nil
}

func (s *InMemoryStorage) SaveAccess(data *osin.AccessData) error {
	//	fmt.Printf("SaveAccess: %s\n", data.AccessToken)
	s.access[data.AccessToken] = data
	if data.RefreshToken != "" {
		s.refresh[data.RefreshToken] = data.AccessToken
	}
	return nil
}

func (s *InMemoryStorage) LoadAccess(code string) (*osin.AccessData, error) {
	//	fmt.Printf("LoadAccess: %s\n", code)
	if d, ok := s.access[code]; ok {
		return d, nil
	}
	return nil, errors.New("Access not found")
}

func (s *InMemoryStorage) RemoveAccess(code string) error {
	//	fmt.Printf("RemoveAccess: %s\n", code)
	delete(s.access, code)
	return nil
}

func (s *InMemoryStorage) LoadRefresh(code string) (*osin.AccessData, error) {
	//	fmt.Printf("LoadRefresh: %s\n", code)
	if d, ok := s.refresh[code]; ok {
		return s.LoadAccess(d)
	}
	return nil, errors.New("Refresh not found")
}

func (s *InMemoryStorage) RemoveRefresh(code string) error {
	//	fmt.Printf("RemoveRefresh: %s\n", code)
	delete(s.refresh, code)
	return nil
}
