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
      <path d="M10 40V24H4L24 6l10 8.85V9h4v9.55L44 24h-6v16H26.5V28h-5v12Zm3-3h5.5V25h11v12H35V19.95l-11-10-11 10Zm5.5-12h11-11Zm1.25-5.5h8.5q0-1.65-1.275-2.725Q25.7 15.7 24 15.7q-1.7 0-2.975 1.075Q19.75 17.85 19.75 19.5Z" />
    </svg>
  )
}

export default AssignmentAdd
