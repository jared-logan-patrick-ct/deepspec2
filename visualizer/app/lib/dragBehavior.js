import * as d3 from 'd3';

/**
 * Setup drag behavior for function cards
 * @param {d3.Selection} selection - D3 selection of draggable elements
 * @param {Function} onPositionChange - Callback when position changes
 */
export function setupDragBehavior(selection, onPositionChange) {
  const drag = d3
    .drag()
    .on('start', function (event) {
      d3.select(this).raise().classed('dragging', true);
    })
    .on('drag', function (event) {
      const current = d3.select(this);
      const currentTransform = current.attr('transform');

      // Extract current position
      const match = currentTransform.match(/translate\(([^,]+),([^)]+)\)/);
      if (match) {
        const x = parseFloat(match[1]) + event.dx;
        const y = parseFloat(match[2]) + event.dy;

        current.attr('transform', `translate(${x}, ${y})`);

        if (onPositionChange) {
          const id = current.attr('data-function-id');
          onPositionChange(id, { x, y });
        }
      }
    })
    .on('end', function (event) {
      d3.select(this).classed('dragging', false);
    });

  selection.call(drag);

  return drag;
}
