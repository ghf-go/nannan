package es_driver

import "net/http"

type esQeury struct {
	*EsClient
	tbName   string
	must     []map[string]map[string]interface{}
	must_not []map[string]map[string]interface{}
	should   []map[string]map[string]interface{}
	from     int
	size     int
	sort     map[string]map[string]interface{}
}

func newEsQuery(ec *EsClient, tname string) *esQeury {
	return &esQeury{
		EsClient: ec,
		tbName:   tname,
		must:     []map[string]map[string]interface{}{},
		must_not: []map[string]map[string]interface{}{},
		should:   []map[string]map[string]interface{}{},
		from:     0,
		size:     0,
		sort:     map[string]map[string]interface{}{},
	}
}
func (eq *esQeury) Start(start int) *esQeury {
	eq.from = start
	return eq
}
func (eq *esQeury) Size(size int) *esQeury {
	eq.size = size
	return eq
}
func (eq *esQeury) buildMap() map[string]interface{} {
	if eq.size == 0 {
		eq.size = 10
	}
	return map[string]interface{}{
		"from": eq.from,
		"size": eq.size,
		"sort": eq.sort,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must":     eq.must,
				"must_not": eq.must_not,
				"should":   eq.should,
			},
		},
	}
}
func (eq *esQeury) Order(key, sort string) *esQeury {
	eq.sort[key] = map[string]interface{}{
		"order": sort,
	}
	return eq
}
func (eq *esQeury) mustkv(k, key string, val interface{}) *esQeury {
	eq.must = append(eq.must, map[string]map[string]interface{}{
		k: {
			key: val,
		},
	})
	return eq
}
func (eq *esQeury) mustnotkv(k string, key string, val interface{}) *esQeury {
	eq.must_not = append(eq.must_not, map[string]map[string]interface{}{
		k: {
			key: val,
		},
	})
	return eq
}
func (eq *esQeury) shouldtkv(k string, key string, val interface{}) *esQeury {
	eq.should = append(eq.should, map[string]map[string]interface{}{
		k: {
			key: val,
		},
	})
	return eq
}

func (eq *esQeury) MustMatch(key string, val interface{}) *esQeury {
	return eq.mustkv("match", key, val)
}

func (eq *esQeury) MustTerm(key string, val interface{}) *esQeury {
	return eq.mustkv("term", key, val)
}
func (eq *esQeury) MustTerms(key string, args ...interface{}) *esQeury {
	return eq.mustkv("terms", key, args)
}
func (eq *esQeury) MustWildcard(key string, val interface{}) *esQeury {
	return eq.mustkv("wildcard", key, val)
}
func (eq *esQeury) MustPrefix(key string, val interface{}) *esQeury {
	return eq.mustkv("prefix", key, val)
}
func (eq *esQeury) MustText(key string, val interface{}) *esQeury {
	return eq.mustkv("text", key, val)
}
func (eq *esQeury) MustQuery(key string, val interface{}) *esQeury {
	eq.must = append(eq.must, map[string]map[string]interface{}{
		"query_string": {
			"default_field": key,
			"query":         val,
		},
	})
	return eq
}
func (eq *esQeury) MustNotQuery(key string, val interface{}) *esQeury {
	eq.must_not = append(eq.must_not, map[string]map[string]interface{}{
		"query_string": {
			"default_field": key,
			"query":         val,
		},
	})
	return eq
}
func (eq *esQeury) ShouldQuery(key string, val interface{}) *esQeury {
	eq.should = append(eq.should, map[string]map[string]interface{}{
		"query_string": {
			"default_field": key,
			"query":         val,
		},
	})
	return eq
}

func (eq *esQeury) MustRange(key, opStart string, start interface{}, opEnd string, end interface{}) *esQeury {
	data := map[string]interface{}{}
	switch opStart {
	case ">":
		data["gt"] = start
	case ">=":
		data["gte"] = start
	}
	switch opEnd {
	case "<":
		data["lt"] = end
	case "<=":
		data["lte"] = end
	}
	eq.should = append(eq.should, map[string]map[string]interface{}{
		"range": {
			key: data,
		},
	})
	return eq
}
func (eq *esQeury) MustNotRange(key, opStart string, start interface{}, opEnd string, end interface{}) *esQeury {
	data := map[string]interface{}{}
	switch opStart {
	case ">":
		data["gt"] = start
	case ">=":
		data["gte"] = start
	}
	switch opEnd {
	case "<":
		data["lt"] = end
	case "<=":
		data["lte"] = end
	}
	eq.must_not = append(eq.must_not, map[string]map[string]interface{}{
		"range": {
			key: data,
		},
	})
	return eq
}
func (eq *esQeury) ShouldRange(key, opStart string, start interface{}, opEnd string, end interface{}) *esQeury {
	data := map[string]interface{}{}
	switch opStart {
	case ">":
		data["gt"] = start
	case ">=":
		data["gte"] = start
	}
	switch opEnd {
	case "<":
		data["lt"] = end
	case "<=":
		data["lte"] = end
	}
	eq.should = append(eq.should, map[string]map[string]interface{}{
		"range": {
			key: data,
		},
	})
	return eq
}

func (eq *esQeury) MustNotMatch(key string, val interface{}) *esQeury {
	return eq.mustnotkv("match", key, val)
}

func (eq *esQeury) MustNotTerm(key string, val interface{}) *esQeury {
	return eq.mustnotkv("term", key, val)
}
func (eq *esQeury) MustNotWildcard(key string, val interface{}) *esQeury {
	return eq.mustnotkv("wildcard", key, val)
}
func (eq *esQeury) MustNotPrefix(key string, val interface{}) *esQeury {
	return eq.mustnotkv("prefix", key, val)
}
func (eq *esQeury) MustNotText(key string, val interface{}) *esQeury {
	return eq.mustnotkv("text", key, val)
}

func (eq *esQeury) ShouldMatch(key string, val interface{}) *esQeury {
	return eq.shouldtkv("match", key, val)
}

func (eq *esQeury) ShouldTerm(key string, val interface{}) *esQeury {
	return eq.shouldtkv("term", key, val)
}
func (eq *esQeury) ShouldWildcard(key string, val interface{}) *esQeury {
	return eq.shouldtkv("wildcard", key, val)
}
func (eq *esQeury) ShouldPrefix(key string, val interface{}) *esQeury {
	return eq.shouldtkv("prefix", key, val)
}
func (eq *esQeury) ShouldText(key string, val interface{}) *esQeury {
	return eq.shouldtkv("text", key, val)
}
func (eq *esQeury) Missing(key string) *esQeury {
	return eq.mustnotkv("exists", "field", key)
}
func (eq *esQeury) Query(obj interface{}) error {
	ret := &esSearchResponse{}
	e := eq.do(http.MethodPost, eq.getHost()+eq.tbName+"/_search", eq.buildMap(), ret)
	if e != nil {
		return e
	}
	return ret.saveObj(obj)
}
