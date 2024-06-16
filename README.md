# Run

## Goroutine

```
go run goroutine/main.go
```

## No goroutine

```
go run nogoroutine/main.go
```

# Results

| manner\receiver | value       | pointer     |
| --------------- | ----------- | ----------- |
| goroutine       | 11.804791ms | 22.644916ms |
| no goroutine    | 43.985708ms | 62.034167ms |
