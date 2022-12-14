interface Props {
  color?: string
  onClick?: () => void
  className?: string
}

const AssignmentAdd = (props: Props) => {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 48 48"
      height="24"
      width="24"
      fill={props.color}
      onClick={props.onClick}
      className={props.className}
    >
      <path d="M13.65 40 4 30.35l9.65-9.65 2.1 2.1-6.05 6.05h15.8v3H9.7l6.05 6.05Zm20.7-12.7-2.1-2.1 6.05-6.05H22.5v-3h15.8l-6.05-6.05 2.1-2.1L44 17.65Z" />
    </svg>
  )
}

export default AssignmentAdd
