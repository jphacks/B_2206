interface Props {
  width?: string
  height?: string
  color?: string
  onClick?: () => void
  className?: string
}

const Domain = (props: Props) => {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 48 48"
      height={props.height ? props.height : '40'}
      width={props.width ? props.width : '40'}
      fill={props.color}
      onClick={props.onClick}
      className={props.className}
    >
      <path d="M4 42V6h19.5v8.25H44V42Zm3-3h5.25v-5.25H7Zm0-8.25h5.25V25.5H7Zm0-8.25h5.25v-5.25H7Zm0-8.25h5.25V9H7ZM15.25 39h5.25v-5.25h-5.25Zm0-8.25h5.25V25.5h-5.25Zm0-8.25h5.25v-5.25h-5.25Zm0-8.25h5.25V9h-5.25ZM23.5 39H41V17.25H23.5v5.25h4v3h-4v5.25h4v3h-4Zm9.25-13.5v-3h3v3Zm0 8.25v-3h3v3Z" />
    </svg>
  )
}

export default Domain
