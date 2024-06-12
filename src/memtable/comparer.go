package memtable

type Comparer interface {
	isLessThanOrEqual(k1, k2 string) bool
	isLessThan(k1, k2 string) bool
}

type StringComparer struct{}

func (sc StringComparer) isLessThanOrEqual(k1, k2 string) bool {
	if k1 == NEG_INF || k2 == POS_INF {
		return true
	}

	if k1 == POS_INF || k2 == NEG_INF {
		return false
	}

	return k1 <= k2
}

func (sc StringComparer) isLessThan(k1, k2 string) bool {
	if k1 == NEG_INF || k2 == POS_INF {
		return true
	}

	if k1 == POS_INF || k2 == NEG_INF {
		return false
	}

	return k1 < k2
}
