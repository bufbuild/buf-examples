go run ./cmd/topic \
    --topic email-updated \
    --topic-config=bufstream.archive.iceberg.catalog=local-rest-catalog \
    --topic-config=bufstream.archive.iceberg.table=bufstream.email_updated \
    --topic-config=bufstream.archive.kind=ICEBERG
