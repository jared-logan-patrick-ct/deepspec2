'use client';

export default function FunctionCard({ func, position, onDragStart, onDrag, onDragEnd }) {
  return (
    <g
      className="function-card"
      transform={`translate(${position.x}, ${position.y})`}
      data-function-id={func.id}
    >
      {/* Card background */}
      <rect
        width="280"
        height="auto"
        rx="8"
        className="fill-bg-medium stroke-border-medium cursor-move"
        strokeWidth="2"
      />

      {/* Header section */}
      <g className="function-header">
        {/* Visibility badge */}
        <rect
          x="12"
          y="12"
          width="60"
          height="20"
          rx="4"
          className={`fill-ct-${func.visibility === 'public' ? 'green' : 'purple'}`}
        />
        <text
          x="42"
          y="26"
          textAnchor="middle"
          className="fill-text-white text-xs font-medium"
        >
          {func.visibility}
        </text>

        {/* Function name */}
        <text
          x="12"
          y="50"
          className="fill-ct-yellow text-lg font-bold"
        >
          {func.name}
        </text>
      </g>

      {/* Description */}
      <text
        x="12"
        y="75"
        className="fill-text-gray text-sm"
      >
        {func.description.length > 35
          ? func.description.substring(0, 35) + '...'
          : func.description
        }
      </text>

      {/* Parameters section */}
      {func.signature && func.signature.parameters && func.signature.parameters.length > 0 && (
        <g className="parameters" transform="translate(0, 90)">
          <text x="12" y="0" className="fill-text-white text-xs font-semibold">
            Parameters:
          </text>
          {func.signature.parameters.slice(0, 3).map((param, idx) => (
            <text
              key={param.name}
              x="12"
              y={20 + idx * 18}
              className="fill-text-gray text-xs font-mono"
            >
              {param.name}: {param.type}
            </text>
          ))}
          {func.signature.parameters.length > 3 && (
            <text
              x="12"
              y={20 + 3 * 18}
              className="fill-text-gray text-xs italic"
            >
              +{func.signature.parameters.length - 3} more...
            </text>
          )}
        </g>
      )}

      {/* Returns section */}
      {func.signature && func.signature.returns && func.signature.returns.length > 0 && (
        <g className="returns" transform="translate(0, 180)">
          <text x="12" y="0" className="fill-text-white text-xs font-semibold">
            Returns:
          </text>
          {func.signature.returns.slice(0, 2).map((ret, idx) => (
            <text
              key={idx}
              x="12"
              y={20 + idx * 18}
              className="fill-text-gray text-xs font-mono"
            >
              {ret.type}
            </text>
          ))}
        </g>
      )}

      {/* Source location footer */}
      <text
        x="12"
        y="240"
        className="fill-text-gray text-xs"
      >
        {func.sourceLocation.file}:{func.sourceLocation.startLine}
      </text>
    </g>
  );
}
