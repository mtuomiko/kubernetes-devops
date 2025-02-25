/// <reference types="vite/client" />

// https://vite.dev/guide/env-and-mode#intellisense-for-typescript

interface ImportMetaEnv {
  readonly VITE_TODO_API_URL: string

}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
