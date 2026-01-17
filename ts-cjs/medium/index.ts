/*
    Medium YINI Demo - TypeScript (CommonJS)

    Demonstrates three ways to parse a YINI file:
    - Variant A: Default parsing
    - Variant B: Strict mode
    - Variant C: Options object + metadata

    This example shows how to:
    - Load and parse settings.yini
    - Read nested values
    - Work with arrays and inline objects
    - Access sections with spaces in their names
    - Inspect both the JS object and JSON output
*/

import YINI = require("yini-parser")

const CONFIG_FILE = "settings.yini"

try {
    console.log("*** Variant A: Default parsing ***\n")

    const config = YINI.parseFile(CONFIG_FILE)

    console.log("=== Application Settings ===")
    console.log("Service ID   :", config.Settings.serviceId)
    console.log("Release      :", config.Settings.release)
    console.log("Debug Mode   :", config.Settings.debugMode)
    console.log("Tagline      :", config.Settings.tagline)
    console.log()

    console.log("=== Network ===")
    console.log("Bind Address :", config.Settings.Network.bindAddress)
    console.log("Bind Port    :", config.Settings.Network.bindPort)
    console.log(
        "Allowed Origins:",
        config.Settings.Network.allowedOrigins.join(", ")
    )
    console.log()

    console.log("=== Capabilities ===")
    console.log("Search enabled:", config.Settings.Capabilities.enableSearch)
    console.log(
        "Experimental features:",
        config.Settings.Capabilities.experimental
    )
    console.log()

    console.log("=== Pools (inline object) ===")
    console.log("Min units :", config.Settings.Capabilities.pools.min_units)
    console.log("Max units :", config.Settings.Capabilities.pools.max_units)
    console.log("Size MB   :", config.Settings.Capabilities.pools.size_mb)
    console.log("Timeout   :", config.Settings.Capabilities.pools.timeout_sec)
    console.log()

    console.log("=== Database ===")
    console.log("DB Host :", config.Settings["DB Config"].host)
    console.log("SSL     :", config.Settings["DB Config"].ssl)
    console.log("User    :", config.Settings["DB Config"].Security.username)
    console.log()

    console.log("--- Parsed Object ---")
    console.dir(config, { depth: 5 })

    console.log("\n--- As JSON ---")
    console.log(JSON.stringify(config, null, 2))
} catch (err) {
    console.error("Error parsing " + CONFIG_FILE + ":", (<any>err)?.message)
}

try {
    console.log("*** Variant B: Strict mode ***\n")

    const strictConfig = YINI.parseFile(CONFIG_FILE, true)

    console.log("Strict mode loaded OK")
    console.log()
    console.log("Service ID:", strictConfig.Settings.serviceId)
    console.log("Bind Port :", strictConfig.Settings.Network.bindPort)
    console.log()
} catch (err) {
    console.error("Error parsing " + CONFIG_FILE + ":", (<any>err)?.message)
}

try {
    console.log("*** Variant C: Options object + metadata ***\n")

    const result = YINI.parseFile(CONFIG_FILE, {
        strictMode: false,
        includeMetadata: true,
        failLevel: "warnings-and-errors"
    })

    console.log("Note: When includeMetadata = true, parseFile returns { result, meta }\n")

    console.log("Parsed result:", result.result)
    console.log("Parser metadata:", result.meta)
} catch (err) {
    console.error("Error parsing " + CONFIG_FILE + ":", (<any>err)?.message)
}
