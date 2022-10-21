interface Props {
  width?: string
  height?: string
  color?: string
  onClick?: () => void
  className?: string
}

const WebSite = (props: Props) => {
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
      <path d="m12.4 35.7-2.1-2.1L30.9 13H12v-3h24v24h-3V15.1Z" />
    </svg>
  )
}

export default WebSite
