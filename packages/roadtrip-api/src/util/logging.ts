// stubs for integration with cloud logging

const logInfo = (message: string) => {
  console.log(message);
};

const logError = (message: string) => {
  console.error(message);
};

const throwError = (e: Error) => {
  logError(e.message);
  throw e;
};

export { logInfo, logError, throwError };
