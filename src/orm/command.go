package orm

type baseCommand interface {
    query(sql string, args map[string]string)
}