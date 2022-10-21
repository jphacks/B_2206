interface Props {
  content?: string
  children?: React.ReactNode
}

const Tooltip = (props: Props) => {
  return (
    <>
      <span className="group relative">
        <span
          className={
            "pointer-events-none absolute -top-12 left-1/2 -translate-x-1/2 whitespace-nowrap rounded bg-black px-2 py-1 text-white opacity-0 transition before:absolute before:left-1/2 before:top-full before:-translate-x-1/2 before:border-4 before:border-transparent before:border-t-black before:content-[''] group-hover:opacity-100"
          }
        >
          {props.content}
        </span>
        {props.children}
      </span>
    </>
  )
}

export default Tooltip
