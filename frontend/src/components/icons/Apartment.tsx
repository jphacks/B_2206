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
      <path d="M6 42V14.25h8.25V6h19.5v16.5H42V42H26.65v-8.25h-5.3V42Zm3-3h5.25v-5.25H9Zm0-8.25h5.25V25.5H9Zm0-8.25h5.25v-5.25H9Zm8.25 8.25h5.25V25.5h-5.25Zm0-8.25h5.25v-5.25h-5.25Zm0-8.25h5.25V9h-5.25Zm8.25 16.5h5.25V25.5H25.5Zm0-8.25h5.25v-5.25H25.5Zm0-8.25h5.25V9H25.5ZM33.75 39H39v-5.25h-5.25Zm0-8.25H39V25.5h-5.25Z" />{' '}
    </svg>
  )
}

export default AssignmentAdd
