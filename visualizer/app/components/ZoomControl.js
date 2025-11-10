'use client';

import { MdCenterFocusStrong } from 'react-icons/md';

export default function ZoomControl({ zoom, onZoomChange, onReset, onRecenter }) {
  const handleSliderChange = (e) => {
    onZoomChange(parseFloat(e.target.value));
  };

  const handleIncrement = () => {
    onZoomChange(Math.min(10, zoom + 0.1));
  };

  const handleDecrement = () => {
    onZoomChange(Math.max(0.1, zoom - 0.1));
  };

  return (
    <div className="fixed bottom-4 right-4 bg-bg-medium p-4 rounded-lg shadow-lg flex flex-col gap-3 min-w-[200px]">
      <div className="flex items-center justify-between gap-3">
        <button
          onClick={handleDecrement}
          className="w-8 h-8 bg-bg-light hover:bg-bg-dark rounded text-text-white font-bold"
        >
          -
        </button>
        <span className="text-text-white font-mono text-sm">
          {Math.round(zoom * 100)}%
        </span>
        <button
          onClick={handleIncrement}
          className="w-8 h-8 bg-bg-light hover:bg-bg-dark rounded text-text-white font-bold"
        >
          +
        </button>
      </div>
      <input
        type="range"
        min="0.1"
        max="10"
        step="0.1"
        value={zoom}
        onChange={handleSliderChange}
        className="w-full"
      />
      <div className="flex gap-2">
        <button
          onClick={onReset}
          className="flex-1 bg-neutral-700 cursor-pointer hover:bg-neutral-600 text-text-white px-3 py-2 rounded text-sm font-medium"
        >
          Reset
        </button>
        <button
          onClick={onRecenter}
          className="w-9 h-9 bg-neutral-700 cursor-pointer hover:bg-neutral-600 text-text-white rounded flex items-center justify-center"
          title="Recenter view"
        >
          <MdCenterFocusStrong size={18} />
        </button>
      </div>
    </div>
  );
}
