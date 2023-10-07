interface TabContainerInterface {
  className?: string;
  children?: React.ReactNode;
}

const TabContainer: React.FC<TabContainerInterface> = ({ children, className }) => {
  return (
    <div
      className={`flex select-none space-x-1 rounded-lg bg-zinc-100 p-1 dark:bg-zinc-800 ${className}`}
    >
      {children}
    </div>
  );
};

export default TabContainer;
