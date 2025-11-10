'use client';

import { useEffect, useRef } from 'react';
import * as d3 from 'd3';

/**
 * Custom hook for D3 visualizations in React
 * @param {Function} renderFn - Function that receives the D3 selection
 * @param {Array} dependencies - Dependencies array for useEffect
 * @returns {Object} - Ref object to attach to DOM element
 */
export function useD3(renderFn, dependencies = []) {
  const ref = useRef();

  useEffect(() => {
    if (ref.current) {
      renderFn(d3.select(ref.current));
    }
    return () => {};
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, dependencies);

  return ref;
}
