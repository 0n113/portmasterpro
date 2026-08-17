/**
 * spn-removed.ts
 *
 * SPN (Safing Privacy Network) and account login have been removed
 * from portmasterpro. This file acts as a single source of truth for
 * UI components that previously depended on SPN/account state.
 *
 * All feature flags are permanently unlocked.
 * All login/account state is permanently absent.
 */

export const SPN_REMOVED = true;

/** Always false — no account session exists. */
export const isLoggedIn = (): boolean => false;

/** Always false — SPN routing is not available. */
export const isSPNActive = (): boolean => false;

/**
 * Always true — all features are available without a subscription.
 * Replace any `hasFeature(x)` call that previously gated UI elements.
 */
export const hasFeature = (_feature: string): boolean => true;

/**
 * No-op — account login UI has been removed.
 * Call sites that previously opened a login modal should call this instead.
 */
export const openLoginModal = (): void => {
  console.warn('[portmasterpro] openLoginModal() called but account login has been removed.');
};

/**
 * No-op — SPN toggle has been removed.
 * Call sites that previously toggled SPN should call this instead.
 */
export const toggleSPN = (): void => {
  console.warn('[portmasterpro] toggleSPN() called but SPN has been removed.');
};
