import clsx from "classnames";

interface Props<T> {
  title: string;
  error?: string;
  onChange: (value: T | undefined) => void;
  value?: T;
  variant: readonly T[];
  className?: string;
}

export function GraduatesToggle<T extends string>({
  variant,
  title,
  onChange,
  value,
  error,
  className,
}: Props<T>) {
  const handleClick = (num: T) => {
    onChange(value === num ? undefined : num);
  };

  return (
    <div className={"flex flex-col w-[48%] min-w-[120px] " + className}>
      <span className="mb-1">{title}</span>
      <div className="flex gap-2">
        {variant.map((num, index) => (
          <button
            key={index}
            type="button"
            onClick={() => handleClick(num)}
            className={clsx(
              "flex-1 rounded border px-3 py-2 transition-colors",
              value === num
                ? "bg-blue-600 text-white border-blue-600"
                : "bg-white text-gray-700 border-gray-300 hover:bg-blue-50",
            )}
          >
            {num}
          </button>
        ))}
      </div>
      {error && <span className="mt-1 text-xs text-red-500">{error}</span>}
    </div>
  );
}
