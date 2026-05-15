from yini_parser import load


def main() -> None:
    config = load("config.yini")

    app = config["App"]
    server = app["Server"]

    print("Application")
    print("-----------")
    print(f"Name: {app['name']}")
    print(f"Version: {app['version']}")
    print(f"Debug: {app['debug']}")
    print(f"Features: {', '.join(app['features'])}")

    print()
    print("Server")
    print("------")
    print(f"Host: {server['host']}")
    print(f"Port: {server['port']}")
    print(f"Use TLS: {server['useTLS']}")


if __name__ == "__main__":
    main()
