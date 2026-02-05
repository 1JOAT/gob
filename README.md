# gob Framework (v0.1.0)

`gob` is a professional, high-performance Go framework for building web applications with real-time hot-reloading and built-in MongoDB support.

## Installation

To install the `gob` CLI globally:

```bash
go install github.com/1joat/gob/cmd/gob@latest
```

## Usage

### Create a new project
By default, `gob` scaffolds projects with MongoDB support.
```bash
gob new myapp
```

#### Build the project
```bash
gob build
```

#### Start development mode
```bash
gob dev
```

## Roadmap

- [ ] Built-in Middleware Support (Logger, Recovery, CORS)
- [ ] Redis Caching Integration
- [ ] Advanced CLI Generators (Models, Controllers, Services)
- [ ] Standardized JSON Error Responses
- [ ] Automated Documentation Generation

---

## Contributing

Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details on how to contribute to the project.

## License

MIT - see [LICENSE](LICENSE) for details.
