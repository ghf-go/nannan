package db

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
func (eq *esQeury) mustkv(k, key, val string) *esQeury {
	eq.must = append(eq.must, map[string]map[string]interface{}{
		k: {
			key: val,
		},
	})
	return eq
}
func (eq *esQeury) mustnotkv(k, key, val string) *esQeury {
	eq.must_not = append(eq.must_not, map[string]map[string]interface{}{
		k: {
			key: val,
		},
	})
	return eq
}
func (eq *esQeury) shouldtkv(k, key, val string) *esQeury {
	eq.should = append(eq.should, map[string]map[string]interface{}{
		k: {
			key: val,
		},
	})
	return eq
}

func (eq *esQeury) MustMatch(key, val string) *esQeury {
	return eq.mustkv("match", key, val)
}

func (eq *esQeury) MustTerm(key, val string) *esQeury {
	return eq.mustkv("term", key, val)
}
func (eq *esQeury) MustWildcard(key, val string) *esQeury {
	return eq.mustkv("wildcard", key, val)
}
func (eq *esQeury) MustPrefix(key, val string) *esQeury {
	return eq.mustkv("prefix", key, val)
}
func (eq *esQeury) MustText(key, val string) *esQeury {
	return eq.mustkv("text", key, val)
}
func (eq *esQeury) MustQuery(key, val string) *esQeury {
	eq.must = append(eq.must, map[string]map[string]interface{}{
		"query_string": {
			"default_field": key,
			"query":         val,
		},
	})
	return eq
}
func (eq *esQeury) MustNotQuery(key, val string) *esQeury {
	eq.must_not = append(eq.must_not, map[string]map[string]interface{}{
		"query_string": {
			"default_field": key,
			"query":         val,
		},
	})
	return eq
}
func (eq *esQeury) ShouldQuery(key, val string) *esQeury {
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

func (eq *esQeury) MustNotMatch(key, val string) *esQeury {
	return eq.mustnotkv("match", key, val)
}

func (eq *esQeury) MustNotTerm(key, val string) *esQeury {
	return eq.mustnotkv("term", key, val)
}
func (eq *esQeury) MustNotWildcard(key, val string) *esQeury {
	return eq.mustnotkv("wildcard", key, val)
}
func (eq *esQeury) MustNotPrefix(key, val string) *esQeury {
	return eq.mustnotkv("prefix", key, val)
}
func (eq *esQeury) MustNotText(key, val string) *esQeury {
	return eq.mustnotkv("text", key, val)
}

func (eq *esQeury) ShouldMatch(key, val string) *esQeury {
	return eq.shouldtkv("match", key, val)
}

func (eq *esQeury) ShouldTerm(key, val string) *esQeury {
	return eq.shouldtkv("term", key, val)
}
func (eq *esQeury) ShouldWildcard(key, val string) *esQeury {
	return eq.shouldtkv("wildcard", key, val)
}
func (eq *esQeury) ShouldPrefix(key, val string) *esQeury {
	return eq.shouldtkv("prefix", key, val)
}
func (eq *esQeury) ShouldText(key, val string) *esQeury {
	return eq.shouldtkv("text", key, val)
}
func (eq *esQeury) Missing(key string) *esQeury {
	return eq.mustnotkv("exists", "field", key)
}
func (eq *esQeury) Query() (*EsResponse, error) {
	return eq.do(http.MethodPost, eq.getHost()+eq.tbName+"/_search", eq.buildMap(), nil)
}
