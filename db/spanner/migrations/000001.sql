CREATE TABLE Todos (
    Id STRING(36) NOT NULL,
    Title STRING(255) NOT NULL,
    Description STRING(MAX),
    Priority INT64 NOT NULL,
    Completed BOOL NOT NULL,
    DueTime TIMESTAMP,
    CreatedAt TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true),
    UpdatedAt TIMESTAMP OPTIONS (allow_commit_timestamp=true),
) PRIMARY KEY (Id);
