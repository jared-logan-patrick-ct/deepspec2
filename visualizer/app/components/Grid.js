import { useState, useEffect, useRef } from 'react';
import * as d3 from 'd3';
import { useD3 } from '../hooks/useD3';
import { setupInfiniteCanvas } from '../lib/d3Canvas';
import { setupDragBehavior } from '../lib/dragBehavior';
import { renderFunctionCard } from '../lib/renderFunctionCard';
import { sampleFunctions } from '../data/sampleFunctions';
import ZoomControl from './ZoomControl';

export default function Grid() {
  const gridSize = 20;
  const [zoom, setZoom] = useState(1);
  const [zoomBehavior, setZoomBehavior] = useState(null);
  const [contentGroup, setContentGroup] = useState(null);
  const [functions, setFunctions] = useState(sampleFunctions);
  const functionsRef = useRef(functions);

  // Keep ref in sync with state
  useEffect(() => {
    functionsRef.current = functions;
  }, [functions]);

  const svgRef = useD3(
    (svg) => {
      const { zoom: zoomBehavior, contentGroup } = setupInfiniteCanvas(svg, gridSize, setZoom);
      setZoomBehavior(() => zoomBehavior);
      setContentGroup(contentGroup);
    },
    [gridSize]
  );

  // Render function cards when content group or functions change
  useEffect(() => {
    if (!contentGroup || functions.length === 0) return;

    // Clear existing cards and render new ones
    contentGroup.selectAll('.function-card').remove();

    functions.forEach((func) => {
      renderFunctionCard(contentGroup, func);
    });

    // Setup drag behavior - use ref to avoid re-render loop
    setupDragBehavior(contentGroup.selectAll('.function-card'), (id, position) => {
      // Update position directly in DOM without triggering re-render
      const card = contentGroup.select(`[data-function-id="${id}"]`);
      if (!card.empty()) {
        card.attr('transform', `translate(${position.x}, ${position.y})`);
      }

      // Update state ref without causing re-render
      functionsRef.current = functionsRef.current.map(f =>
        f.id === id ? { ...f, position } : f
      );
    });
  }, [contentGroup, functions.length]);

  const handleZoomChange = (newZoom) => {
    if (zoomBehavior && svgRef.current) {
      const svg = d3.select(svgRef.current);
      zoomBehavior.scaleTo(svg, newZoom);
    }
  };

  const handleReset = () => {
    handleZoomChange(1);
  };

  const handleRecenter = () => {
    if (zoomBehavior && svgRef.current) {
      const svg = d3.select(svgRef.current);
      const node = svgRef.current;
      const width = node.clientWidth;
      const height = node.clientHeight;

      // Reset to initial position (centered)
      svg.transition()
        .duration(300)
        .call(zoomBehavior.transform, d3.zoomIdentity.translate(width / 2, height / 2).scale(zoom));
    }
  };

  return (
    <>
      <svg
        ref={svgRef}
        className="w-full h-full bg-bg-darkest cursor-default"
      />
      <ZoomControl
        zoom={zoom}
        onZoomChange={handleZoomChange}
        onReset={handleReset}
        onRecenter={handleRecenter}
      />
    </>
  );
}
