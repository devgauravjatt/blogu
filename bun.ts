import { $ } from "bun";

await $`./blogu.exe build && cd build && live-server`.text();
