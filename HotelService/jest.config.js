const { createDefaultPreset } = require("ts-jest");

const tsJestTransformCfg = createDefaultPreset().transform;

/** @type {import("jest").Config} **/
module.exports = {
  preset: "ts-jest",
  testEnvironment: "node",
  transform: {
    ...tsJestTransformCfg,
  },
  roots: ["<rootDir>/src", "<rootDir>/tests"],
  testMatch: ["**/?(*.)+(spec|test).ts"],
  setupFilesAfterEnv: ["<rootDir>/tests/setup.ts"],
  collectCoverageFrom: [
    "src/**/*.ts",
    "!src/server.ts", // avoid starting server/workers
  ],
  coverageDirectory: "coverage",
  globals: {
    "ts-jest": {
      tsconfig: "<rootDir>/tsconfig.jest.json",
    },
  },
};