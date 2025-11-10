import * as d3 from 'd3';
import { colors } from '../config/colors';

/**
 * Setup infinite canvas with pan behavior
 * @param {d3.Selection} svg - D3 selection of SVG element
 * @param {number} gridSize - Size of grid cells in pixels
 * @param {Function} onZoomChange - Callback for zoom level changes
 */
export function setupInfiniteCanvas(svg, gridSize, onZoomChange) {
  const node = svg.node();
  const width = node.clientWidth;
  const height = node.clientHeight;

  // Clear existing content
  svg.selectAll('*').remove();

  // Create defs for grid pattern
  const defs = svg.append('defs');
  const pattern = defs
    .append('pattern')
    .attr('id', 'grid')
    .attr('width', gridSize)
    .attr('height', gridSize)
    .attr('patternUnits', 'userSpaceOnUse');

  pattern
    .append('rect')
    .attr('width', gridSize)
    .attr('height', gridSize)
    .attr('fill', colors.bgDark);

  pattern
    .append('path')
    .attr('d', `M ${gridSize} 0 L 0 0 0 ${gridSize}`)
    .attr('fill', 'none')
    .attr('stroke', colors.borderMedium)
    .attr('stroke-width', 1);

  // Create main group for panning
  const g = svg.append('g');

  // Add truly infinite grid background
  g.append('rect')
    .attr('x', -1000000)
    .attr('y', -1000000)
    .attr('width', 2000000)
    .attr('height', 2000000)
    .attr('fill', 'url(#grid)');

  // Content group for future draggable items
  const contentGroup = g.append('g').attr('class', 'content-group');

  // Setup zoom and pan behavior
  const zoom = d3
    .zoom()
    .scaleExtent([0.1, 10])
    .filter((event) => {
      // Allow wheel events for zoom
      if (event.type === 'wheel') return true;
      // Only allow pan with middle mouse button (button 1)
      if (event.type === 'mousedown') {
        const isMiddleButton = event.button === 1;
        if (isMiddleButton) {
          svg.style('cursor', 'grabbing');
        }
        return isMiddleButton;
      }
      // Allow drag if middle button started the interaction
      if (event.type === 'mousemove' && event.buttons === 4) {
        return true;
      }
      return false;
    })
    .on('zoom', (event) => {
      g.attr('transform', event.transform);
      if (onZoomChange) {
        onZoomChange(event.transform.k);
      }
    })
    .on('end', (event) => {
      if (event.sourceEvent?.type === 'mouseup') {
        svg.style('cursor', 'default');
      }
    });

  // Apply zoom behavior to svg
  svg.call(zoom);

  // Center the view initially
  svg.call(zoom.transform, d3.zoomIdentity.translate(width / 2, height / 2));

  return {
    svg,
    mainGroup: g,
    contentGroup,
    zoom,
  };
}
