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
| goroutine       | 11.804791ms | 12.798125ms |
| no goroutine    | 43.985708ms | 19.081292ms |
