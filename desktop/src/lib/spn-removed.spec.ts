/**
 * spn-removed.spec.ts
 *
 * Unit tests for the SPN/account removal shim.
 * These run with any standard TypeScript test runner (Vitest / Jest).
 */

import { describe, it, expect } from 'vitest';
import {
  SPN_REMOVED,
  isLoggedIn,
  isSPNActive,
  hasFeature,
} from './spn-removed';

describe('SPN/account removal shim', () => {
  it('SPN_REMOVED flag is true', () => {
    expect(SPN_REMOVED).toBe(true);
  });

  it('isLoggedIn() always returns false', () => {
    expect(isLoggedIn()).toBe(false);
  });

  it('isSPNActive() always returns false', () => {
    expect(isSPNActive()).toBe(false);
  });

  it('hasFeature() always returns true for any feature string', () => {
    const features = [
      'bandwidth-visibility',
      'network-history',
      'filter-lists',
      'spn',
      'custom-lists',
      'some-future-feature',
    ];
    for (const f of features) {
      expect(hasFeature(f)).toBe(true);
    }
  });

  it('hasFeature() returns true even for empty string', () => {
    expect(hasFeature('')).toBe(true);
  });
});
