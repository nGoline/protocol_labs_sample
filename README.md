# Quanto Custa o Bitcoin

Website that keeps track of the median price of Bitcoin on the main exchanges in Brazil
and display a graph where each candle is represented by a Bitocin block.

## Running on development

To compile the executables run:

```bash
make install && make build
```

To start the exchange workers run:

```bash
./bin/exchangeworker <exchangename>
```

In both cases you can run `./bin/<workername> help` to list the available commands.
