interface Props {
  width?: number
  height?: number
  color?: string
  onClick?: () => void
  className?: string
}

const AssignmentAdd = (props: Props) => {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 48 48"
      height="40"
      width="40"
      fill={props.color}
      onClick={props.onClick}
      className={props.className}
    >
      <path d="M9 42q-1.25 0-2.125-.875T6 39V9q0-1.25.875-2.125T9 6h10.25q.25-1.75 1.6-2.875T24 2q1.8 0 3.15 1.125Q28.5 4.25 28.75 6H39q1.25 0 2.125.875T42 9v16.45q-.7-.4-1.475-.65-.775-.25-1.525-.4V9H9v30h15.45q.2.8.45 1.55t.6 1.45Zm0-5.35V39 9v15.4-.2 12.45ZM14 34h10.55q.2-.8.45-1.55t.65-1.45H14Zm0-8.5h17.2q.7-.35 1.35-.575.65-.225 1.45-.425v-2H14Zm0-8.5h20v-3H14Zm10-8.85q.7 0 1.225-.525.525-.525.525-1.225 0-.7-.525-1.225Q24.7 4.65 24 4.65q-.7 0-1.225.525-.525.525-.525 1.225 0 .7.525 1.225.525.525 1.225.525Zm12.65 37.8q-3.9 0-6.65-2.775-2.75-2.775-2.75-6.575 0-3.9 2.75-6.675t6.65-2.775q3.85 0 6.625 2.775T46.05 36.6q0 3.8-2.775 6.575Q40.5 45.95 36.65 45.95Zm-.75-3h1.65v-5.5h5.5V35.8h-5.5v-5.5H35.9v5.5h-5.5v1.65h5.5Z" />
    </svg>
  )
}

export default AssignmentAdd
