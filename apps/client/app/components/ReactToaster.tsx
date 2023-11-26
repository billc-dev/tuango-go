import toast, { ToastBar, Toaster } from "react-hot-toast";

const ReactToaster = () => {
  return (
    <Toaster
      containerClassName="mx-auto max-w-3xl"
      toastOptions={{
        duration: 5000,
        position: "top-right",
        className: "flex items-center dark:bg-zinc-600 dark:text-zinc-50",
      }}
    >
      {(t) => (
        <ToastBar toast={t}>
          {({ icon, message }) => (
            <div style={{ ...t.style }} className={t.className} onClick={() => toast.dismiss(t.id)}>
              {icon}
              {message}
            </div>
          )}
        </ToastBar>
      )}
    </Toaster>
  );
};

export default ReactToaster;
