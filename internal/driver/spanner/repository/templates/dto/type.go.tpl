{{- $short := (shortname .Name "err" "res" "sqlstr" "db" "YOLog") -}}
{{- $table := (.Table.TableName) -}}
// {{ .Name }}Dto represents a row from '{{ $table }}'.
type {{ .Name }}Dto struct {
{{- range .Fields }}
{{- if eq (.Col.DataType) (.Col.ColumnName) }}
	{{ .Name }} string `spanner:"{{ .Col.ColumnName }}" json:"{{ .Col.ColumnName }}"` // {{ .Col.ColumnName }} enum
{{- else }}
	{{ .Name }} {{ .Type }} `spanner:"{{ .Col.ColumnName }}" json:"{{ .Col.ColumnName }}"` // {{ .Col.ColumnName }}
{{- end }}
{{- end }}
}

func ({{ $short }} *{{ .Name }}Dto) ToEntity() *domain.{{ .Name }} {
	return &domain.{{ .Name }}{
{{- range .Fields }}
	{{- if eq .Type "time.Time" }}
		{{ .Name }}: {{ $short }}.{{ .Name }}.In(time.FixedZone("JST", 9*60*60)),
	{{- else if eq .Type "spanner.NullTime" }}
        {{ .Name }}: {{ $short }}.{{ .Name }}.Time.In(time.FixedZone("JST", 9*60*60)),
	{{- else if eq .Type "spanner.NullString" }}
        {{ .Name }}: {{ $short }}.{{ .Name }}.String(),
    {{- else if .CustomType }}
        {{ .Name }}: {{ retype .CustomType }}({{ $short }}.{{ .Name }}),
    {{- else }}
		{{ .Name }}: {{ $short }}.{{ .Name }},
	{{- end }}
{{- end }}
	}
}

{{ if .PrimaryKey }}
func ({{ $short }} *{{ .Name }}Dto) primaryKeys() []string {
     return []string{
{{- range .PrimaryKeyFields }}
		"{{ colname .Col }}",
{{- end }}
	}
}
{{- end }}

func ({{ $short }} *{{ .Name }}Dto) writableColumns() []string {
	return []string{
{{- range .Fields }}
	{{- if not .Col.IsGenerated }}
		"{{ colname .Col }}",
	{{- end }}
{{- end }}
	}
}

func ({{ $short }} *{{ .Name }}Dto) columnsToPtrs(cols []string, customPtrs map[string]interface{}) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		if val, ok := customPtrs[col]; ok {
			ret = append(ret, val)
			continue
		}

		switch col {
{{- range .Fields }}
		case "{{ colname .Col }}":
			ret = append(ret, &{{ $short }}.{{ .Name }})
{{- end }}
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}
	return ret, nil
}

func ({{ $short }} *{{ .Name }}Dto) columnsToValues(cols []string) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		switch col {
{{- range .Fields }}
		case "{{ colname .Col }}":
			{{- if .CustomType }}
			ret = append(ret, {{ .Type }}({{ $short }}.{{ .Name }}))
			{{- else }}
			ret = append(ret, {{ $short }}.{{ .Name }})
			{{- end }}
{{- end }}
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}

	return ret, nil
}

// Insert returns a Mutation to insert a row into a table. If the row already
// exists, the write or transaction fails.
func ({{ $short }} *{{ .Name }}Dto) Insert() *spanner.Mutation {
	values, _ := {{ $short }}.columnsToValues({{ $short }}.writableColumns())
	return spanner.Insert("{{ $table }}", {{ $short }}.writableColumns(), values)
}

// Update returns a Mutation to update a row in a table. If the row does not
// already exist, the write or transaction fails.
func ({{ $short }} *{{ .Name }}Dto) Update() *spanner.Mutation {
	values, _ := {{ $short }}.columnsToValues({{ $short }}.writableColumns())
	return spanner.Update("{{ $table }}", {{ $short }}.writableColumns(), values)
}

// Delete deletes the {{ .Name }} from the database.
func ({{ $short }} *{{ .Name }}Dto) Delete() *spanner.Mutation {
	values, _ := {{ $short }}.columnsToValues({{ $short }}.primaryKeys())
	return spanner.Delete("{{ $table }}", spanner.Key(values))
}
