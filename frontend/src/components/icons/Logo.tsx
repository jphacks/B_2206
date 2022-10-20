interface Props {
  width?: string
  height?: string
}

const Add = (props: Props) => {
  const width = props.width ? props.width : '280'
  const height = props.height ? props.height : '45'
  return (
    <svg
      width={width}
      height={height}
      viewBox="0 0 591 106"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        d="M57 2L7 38H57V81.5C57 86.4706 61.0294 90.5 66 90.5H107L143 103V51L107 38H162V90.5H541.5"
        stroke="black"
        strokeWidth="4"
      />
      <path
        d="M207.972 66.6307C207.835 65.2557 207.25 64.1875 206.216 63.4261C205.182 62.6648 203.778 62.2841 202.006 62.2841C200.801 62.2841 199.784 62.4545 198.955 62.7955C198.125 63.125 197.489 63.5852 197.045 64.1761C196.614 64.767 196.398 65.4375 196.398 66.1875C196.375 66.8125 196.506 67.358 196.79 67.8239C197.085 68.2898 197.489 68.6932 198 69.0341C198.511 69.3636 199.102 69.6534 199.773 69.9034C200.443 70.142 201.159 70.3466 201.92 70.517L205.057 71.267C206.58 71.608 207.977 72.0625 209.25 72.6307C210.523 73.1989 211.625 73.8977 212.557 74.7273C213.489 75.5568 214.21 76.5341 214.722 77.6591C215.244 78.7841 215.511 80.0739 215.523 81.5284C215.511 83.6648 214.966 85.517 213.886 87.0852C212.818 88.642 211.273 89.8523 209.25 90.7159C207.239 91.5682 204.813 91.9943 201.972 91.9943C199.153 91.9943 196.699 91.5625 194.608 90.6989C192.528 89.8352 190.903 88.5568 189.733 86.8636C188.574 85.1591 187.966 83.0511 187.909 80.5398H195.051C195.131 81.7102 195.466 82.6875 196.057 83.4716C196.659 84.2443 197.46 84.8295 198.46 85.2273C199.472 85.6136 200.614 85.8068 201.886 85.8068C203.136 85.8068 204.222 85.625 205.142 85.2614C206.074 84.8977 206.795 84.392 207.307 83.7443C207.818 83.0966 208.074 82.3523 208.074 81.5114C208.074 80.7273 207.841 80.0682 207.375 79.5341C206.92 79 206.25 78.5455 205.364 78.1705C204.489 77.7955 203.415 77.4545 202.142 77.1477L198.341 76.1932C195.398 75.4773 193.074 74.358 191.369 72.8352C189.665 71.3125 188.818 69.2614 188.83 66.6818C188.818 64.5682 189.381 62.7216 190.517 61.142C191.665 59.5625 193.239 58.3295 195.239 57.4432C197.239 56.5568 199.511 56.1136 202.057 56.1136C204.648 56.1136 206.909 56.5568 208.841 57.4432C210.784 58.3295 212.295 59.5625 213.375 61.142C214.455 62.7216 215.011 64.5511 215.045 66.6307H207.972ZM244.69 56.5909H252.07V79.2614C252.07 81.8068 251.462 84.0341 250.247 85.9432C249.042 87.8523 247.354 89.3409 245.184 90.4091C243.014 91.4659 240.485 91.9943 237.599 91.9943C234.701 91.9943 232.167 91.4659 229.997 90.4091C227.826 89.3409 226.139 87.8523 224.934 85.9432C223.729 84.0341 223.127 81.8068 223.127 79.2614V56.5909H230.508V78.6307C230.508 79.9602 230.798 81.142 231.377 82.1761C231.968 83.2102 232.798 84.0227 233.866 84.6136C234.934 85.2045 236.178 85.5 237.599 85.5C239.031 85.5 240.275 85.2045 241.332 84.6136C242.4 84.0227 243.224 83.2102 243.803 82.1761C244.394 81.142 244.69 79.9602 244.69 78.6307V56.5909ZM260.783 56.5909H269.885L279.499 80.0455H279.908L289.521 56.5909H298.624V91.5H291.465V68.7784H291.175L282.141 91.3295H277.266L268.232 68.6932H267.942V91.5H260.783V56.5909ZM314.725 56.5909V91.5H307.345V56.5909H314.725ZM323.438 56.5909H332.54L342.154 80.0455H342.563L352.176 56.5909H361.279V91.5H354.12V68.7784H353.83L344.796 91.3295H339.921L330.887 68.6932H330.597V91.5H323.438V56.5909ZM376.017 91.5H368.108L380.159 56.5909H389.67L401.704 91.5H393.795L385.051 64.5682H384.778L376.017 91.5ZM375.522 77.7784H394.204V83.5398H375.522V77.7784ZM402.951 62.6761V56.5909H431.621V62.6761H420.934V91.5H413.638V62.6761H402.951ZM468.19 68.8125H460.724C460.588 67.8466 460.31 66.9886 459.889 66.2386C459.469 65.4773 458.929 64.8295 458.27 64.2955C457.611 63.7614 456.849 63.3523 455.986 63.0682C455.133 62.7841 454.207 62.642 453.207 62.642C451.4 62.642 449.827 63.0909 448.486 63.9886C447.145 64.875 446.105 66.1705 445.366 67.875C444.628 69.5682 444.258 71.625 444.258 74.0455C444.258 76.5341 444.628 78.625 445.366 80.3182C446.116 82.0114 447.162 83.2898 448.503 84.1534C449.844 85.017 451.395 85.4489 453.156 85.4489C454.145 85.4489 455.06 85.3182 455.9 85.0568C456.753 84.7955 457.508 84.4148 458.167 83.9148C458.827 83.4034 459.372 82.7841 459.804 82.0568C460.247 81.3295 460.554 80.5 460.724 79.5682L468.19 79.6023C467.997 81.2045 467.514 82.75 466.741 84.2386C465.98 85.7159 464.952 87.0398 463.656 88.2102C462.372 89.3693 460.838 90.2898 459.054 90.9716C457.281 91.642 455.275 91.9773 453.037 91.9773C449.923 91.9773 447.139 91.2727 444.685 89.8636C442.241 88.4545 440.31 86.4148 438.889 83.7443C437.48 81.0739 436.775 77.8409 436.775 74.0455C436.775 70.2386 437.491 67 438.923 64.3295C440.355 61.6591 442.298 59.625 444.753 58.2273C447.207 56.8182 449.969 56.1136 453.037 56.1136C455.06 56.1136 456.935 56.3977 458.662 56.9659C460.4 57.5341 461.94 58.3636 463.281 59.4545C464.622 60.5341 465.713 61.858 466.554 63.4261C467.406 64.9943 467.952 66.7898 468.19 68.8125ZM476.123 91.5V56.5909H483.504V70.9943H498.486V56.5909H505.85V91.5H498.486V77.0795H483.504V91.5H476.123Z"
        fill="black"
      />
    </svg>
  )
}

export default Add
