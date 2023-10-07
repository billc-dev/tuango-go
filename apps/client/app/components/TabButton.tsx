interface TabButtonInterface extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  selected: boolean;
  icon?: JSX.Element;
}

const TabButton: React.FC<TabButtonInterface> = ({ selected, children, icon, ...props }) => {
  return (
    <button
      className={`flex w-full items-center justify-center rounded-lg py-2.5 text-sm font-medium leading-5 transition-all ${
        selected
          ? "bg-zinc-600 text-white shadow-inner dark:bg-zinc-300 dark:text-black dark:outline-none"
          : "hover:bg-zinc-300 dark:text-zinc-100 dark:outline-none dark:hover:bg-zinc-600"
      }`}
      {...props}
    >
      {icon && <div className="mr-1 h-5 w-5">{icon}</div>}
      {children}
    </button>
  );
};

export default TabButton;
