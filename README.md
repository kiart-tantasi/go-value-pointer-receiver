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
| goroutine       | 12.798125ms | 19.081292ms |
| no goroutine    | 43.985708ms | 62.034167ms |
