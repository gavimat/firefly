CREATE TABLE contractapis (
  seq               INTEGER         PRIMARY KEY AUTOINCREMENT,
  id                UUID            NOT NULL,
  interface_id      UUID            NOT NULL,
  ledger            TEXT,
  location          TEXT,
  name              VARCHAR(64)     NOT NULL,
  namespace         VARCHAR(64)     NOT NULL,
  message_id        UUID            NOT NULL
);

CREATE UNIQUE INDEX contractapis_namespace_name ON contractapis(namespace,name);
