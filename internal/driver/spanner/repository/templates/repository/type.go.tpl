{{- $short := (shortname .Name "err" "res" "sqlstr" "db" "YOLog") -}}
{{- $table := (.Table.TableName) -}}

const {{ .Name }}TableName = "{{ $table }}"

// {{ .Name }}Repository provides access to {{ .Name }} rows.
type {{ .Name }}Repository struct {
    client *spanner.Client
}

func (repo *{{ .Name }}Repository) columns() []string {
	return []string{
{{- range .Fields }}
		"{{ colname .Col }}",
{{- end }}
	}
}

func (repo *{{ .Name }}Repository) ToDto(entity *domain.{{ .Name }}) *{{ .Name }}Dto {
    dto := &{{ .Name }}Dto{}
{{- range .Fields }}
	{{- if eq .Type "spanner.NullTime" }}
		if entity.{{ .Name }}.IsZero() {
			dto.{{ .Name }} = spanner.NullTime{}
		} else {
			dto.{{ .Name }} = spanner.NullTime{Time: entity.{{ .Name }}, Valid: true}
		}
	{{- else if eq .Type "spanner.NullString" }}
		if entity.{{ .Name }} == "" {
			dto.{{ .Name }} = spanner.NullString{}
		} else {
			dto.{{ .Name }} = spanner.NullString{StringVal: entity.{{ .Name }}, Valid: true}
		}
	{{- else if .CustomType }}
			dto.{{ .Name }} = {{ retype .Type }}(entity.{{ .Name }})
    {{- else }}
		dto.{{ .Name }} = entity.{{ .Name }}
	{{- end }}
{{- end }}
	return dto
}

// new{{ .Name }}_Decoder returns a decoder which reads a row from *spanner.Row
// into {{ .Name }}. The decoder is not goroutine-safe. Don't use it concurrently.
func (repo *{{ .Name }}Repository) newDecoder(cols []string) func(*spanner.Row) (*{{ .Name }}Dto, error) {
	customPtrs := map[string]interface{}{}

	return func(row *spanner.Row) (*{{ .Name }}Dto, error) {
        var {{ $short }} {{ .Name }}Dto
        ptrs, err := {{ $short }}.columnsToPtrs(cols, customPtrs)
        if err != nil {
            return nil, err
        }

        if err := row.Columns(ptrs...); err != nil {
            return nil, err
        }

		return &{{ $short }}, nil
	}
}

func (repo *{{ .Name }}Repository) Insert(entity *domain.{{ .Name }}) *spanner.Mutation {
	return repo.ToDto(entity).Insert()
}

func (repo *{{ .Name }}Repository) Update(entity *domain.{{ .Name }}) *spanner.Mutation {
	return repo.ToDto(entity).Update()
}

func (repo *{{ .Name }}Repository) Delete(entity *domain.{{ .Name }}) *spanner.Mutation {
	return repo.ToDto(entity).Delete()
}

{{ if ne (fieldnames .Fields $short .PrimaryKeyFields) "" }}
// Find{{ .Name }} gets a {{ .Name }} by primary key
func (repo *{{ .Name }}Repository) Find(ctx context.Context, db domain_repo.YORODB{{ gocustomparamlist .PrimaryKeyFields true true }}) (*domain.{{ .Name }}, error) {
	key := spanner.Key{ {{ gocustomparamlist .PrimaryKeyFields false false }} }
	row, err := db.ReadRow(ctx, "{{ $table }}", key, repo.columns())
	if err != nil {
		return nil, newError("Find{{ .Name }}", "{{ $table }}", err)
	}

	decoder := repo.newDecoder(repo.columns())
	{{ $short }}, err := decoder(row)
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "Find", "{{ $table }}", err)
	}

	return {{ $short }}.ToEntity(), nil
}
{{ end }}

func (repo *{{ .Name }}Repository) FindByStatement(ctx context.Context, db domain_repo.YORODB, stmt spanner.Statement) ([]*domain.{{ .Name }}, error) {
	iter := db.Query(ctx, stmt)
	defer iter.Stop()

    decoder := repo.newDecoder(repo.columns())
	var {{ $short }}s []*domain.{{ .Name }}
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

        {{ $short }}, err := decoder(row)
    	if err != nil {
    		return nil, newErrorWithCode(codes.Internal, "FindByStatement", "{{ $table }}", err)
        }

		{{ $short }}s = append({{ $short }}s, {{ $short }}.ToEntity())
	}

	return {{ $short }}s, nil
}


