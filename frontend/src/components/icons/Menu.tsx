interface Props {
  width?: number
  height?: number
  color?: string
  onClick?: () => void
  className?: string
}

const Menu = (props: Props) => {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      height={props.height ? props.height : '32'}
      width={props.width ? props.width : '32'}
      viewBox="0 0 48 48"
      fill={props.color}
      onClick={props.onClick}
      className={props.className}
    >
      <path d="M6 36v-3h36v3Zm0-10.5v-3h36v3ZM6 15v-3h36v3Z" />
    </svg>
  )
}

export default Menu
