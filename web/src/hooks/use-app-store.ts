import { useSyncExternalStore } from "react";
import { store, type AppState } from "@/lib/store";

/**
 * Subscribe a component to the whole app state (SPEC §16). Derive per-view data
 * with useMemo from this; getState returns a stable reference between changes so
 * useSyncExternalStore bails on Object.is correctly.
 */
export function useAppState(): AppState {
  return useSyncExternalStore(store.subscribe, store.getState, store.getState);
}

/** First allowed workspace root, or "" before the session loads. Used to open
 * directory pickers at a valid location. */
export function useWorkspaceRoot(): string {
  return useAppState().session?.workspaceRoots?.[0] ?? "";
}
