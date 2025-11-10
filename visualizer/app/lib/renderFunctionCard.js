/**
 * Render a function card as SVG elements
 * @param {d3.Selection} parentGroup - Parent D3 selection to append to
 * @param {Object} func - Function data matching FunctionComponent schema
 */
export function renderFunctionCard(parentGroup, func) {
  const cardGroup = parentGroup
    .append('g')
    .attr('class', 'function-card cursor-move')
    .attr('transform', `translate(${func.position.x}, ${func.position.y})`)
    .attr('data-function-id', func.id);

  // Card background
  cardGroup
    .append('rect')
    .attr('width', 280)
    .attr('height', 260)
    .attr('rx', 8)
    .attr('fill', '#1a1a1a')
    .attr('stroke', '#2a2a2a')
    .attr('stroke-width', 2);

  // Visibility badge
  cardGroup
    .append('rect')
    .attr('x', 12)
    .attr('y', 12)
    .attr('width', 60)
    .attr('height', 20)
    .attr('rx', 4)
    .attr('fill', func.visibility === 'public' ? '#13C1C1' : '#6259FE');

  cardGroup
    .append('text')
    .attr('x', 42)
    .attr('y', 26)
    .attr('text-anchor', 'middle')
    .attr('fill', '#ffffff')
    .attr('font-size', '12px')
    .attr('font-weight', '500')
    .text(func.visibility);

  // Function name
  cardGroup
    .append('text')
    .attr('x', 12)
    .attr('y', 50)
    .attr('fill', '#FFC806')
    .attr('font-size', '18px')
    .attr('font-weight', 'bold')
    .text(func.name);

  // Description
  const desc = func.description.length > 35
    ? func.description.substring(0, 35) + '...'
    : func.description;

  cardGroup
    .append('text')
    .attr('x', 12)
    .attr('y', 75)
    .attr('fill', '#9ca3af')
    .attr('font-size', '14px')
    .text(desc);

  // Parameters section
  if (func.signature?.parameters?.length > 0) {
    cardGroup
      .append('text')
      .attr('x', 12)
      .attr('y', 105)
      .attr('fill', '#ffffff')
      .attr('font-size', '12px')
      .attr('font-weight', '600')
      .text('Parameters:');

    func.signature.parameters.slice(0, 3).forEach((param, pidx) => {
      cardGroup
        .append('text')
        .attr('x', 12)
        .attr('y', 125 + pidx * 18)
        .attr('fill', '#9ca3af')
        .attr('font-size', '12px')
        .attr('font-family', 'monospace')
        .text(`${param.name}: ${param.type}`);
    });

    if (func.signature.parameters.length > 3) {
      cardGroup
        .append('text')
        .attr('x', 12)
        .attr('y', 125 + 3 * 18)
        .attr('fill', '#9ca3af')
        .attr('font-size', '12px')
        .attr('font-style', 'italic')
        .text(`+${func.signature.parameters.length - 3} more...`);
    }
  }

  // Returns section
  if (func.signature?.returns?.length > 0) {
    cardGroup
      .append('text')
      .attr('x', 12)
      .attr('y', 195)
      .attr('fill', '#ffffff')
      .attr('font-size', '12px')
      .attr('font-weight', '600')
      .text('Returns:');

    func.signature.returns.slice(0, 2).forEach((ret, ridx) => {
      cardGroup
        .append('text')
        .attr('x', 12)
        .attr('y', 215 + ridx * 18)
        .attr('fill', '#9ca3af')
        .attr('font-size', '12px')
        .attr('font-family', 'monospace')
        .text(ret.type);
    });
  }

  // Source location
  cardGroup
    .append('text')
    .attr('x', 12)
    .attr('y', 248)
    .attr('fill', '#9ca3af')
    .attr('font-size', '12px')
    .text(`${func.sourceLocation.file}:${func.sourceLocation.startLine}`);

  return cardGroup;
}
