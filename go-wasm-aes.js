const loadGoWasmAes = async (wasmUrl) => {
  const go = new Go();

  const result = await WebAssembly.instantiateStreaming(
    fetch(wasmUrl),
    go.importObject
  );

  await go.run(result.instance);
  console.log("Result: ", encrypt("p", "s", "t"));
};
