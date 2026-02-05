/**
 * Type-safe API client for VitalStack backend
 * Generated types come from openapi.yaml via openapi-typescript
 *
 * In development, Vite proxies /api/* to the backend (see vite.config.ts)
 * In production, set PUBLIC_API_URL environment variable at runtime
 */
import createClient from "openapi-fetch";
import { env } from "$env/dynamic/public";
import type { paths } from "./schema";

// Use PUBLIC_API_URL from runtime environment, fallback to empty (Vite proxy)
// $env/dynamic/public reads from the server's environment at RUNTIME, not build time
const apiBaseUrl = env.PUBLIC_API_URL || "";

export const api = createClient<paths>({
  baseUrl: apiBaseUrl
});

// Re-export types for convenience
export type { paths, components, operations } from "./schema";
