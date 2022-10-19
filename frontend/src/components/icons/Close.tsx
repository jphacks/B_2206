interface Props {
  width?: string
  height?: string
  color?: string
  onClick?: () => void
}

const Close = (props: Props) => {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      height={props.height ? props.height : '40'}
      width={props.width ? props.width : '40'}
      viewBox="0 0 48 48"
      fill={props.color}
      onClick={props.onClick}
    >
      <path d="m12.45 37.65-2.1-2.1L21.9 24 10.35 12.45l2.1-2.1L24 21.9l11.55-11.55 2.1 2.1L26.1 24l11.55 11.55-2.1 2.1L24 26.1Z" />
    </svg>
  )
}

export default Close
