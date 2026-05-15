from yini_parser import load


def main() -> None:
    settings = load("settings.yini")

    app = settings["App"]
    runtime = app["Runtime"]
    server = app["Server"]
    database = app["Database"]
    features = app["Features"]
    email = app["Email"]

    print("Application Settings")
    print("====================")
    print(f"Name: {app['name']}")
    print(f"Environment: {app['environment']}")
    print(f"Debug mode: {app['debug']}")

    print()
    print("Runtime")
    print("-------")
    print(f"Log level: {runtime['logLevel']}")
    print(f"Timezone: {runtime['timezone']}")
    print(f"Workers: {runtime['workers']}")

    print()
    print("Server")
    print("------")
    print(f"Host: {server['host']}")
    print(f"Port: {server['port']}")
    print(f"Public URL: {server['publicUrl']}")

    print()
    print("Database")
    print("--------")
    print(f"Type: {database['type']}")
    print(f"Host: {database['host']}")
    print(f"Port: {database['port']}")
    print(f"Database: {database['database']}")
    print(f"Pool min: {database['pool']['min']}")
    print(f"Pool max: {database['pool']['max']}")

    print()
    print("Enabled Features")
    print("----------------")
    for feature in features["enabled"]:
        print(f"- {feature}")

    print()
    print("Email")
    print("-----")
    print(f"Provider: {email['provider']}")
    print(f"From address: {email['fromAddress']}")
    print(f"Retry count: {email['retryCount']}")


if __name__ == "__main__":
    main()
