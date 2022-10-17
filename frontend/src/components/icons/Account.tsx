interface Props {
  width?: number
  height?: number
  color?: string
  onClick?: () => void
}

const Add = (props: Props) => {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      height="40"
      width="40"
      fill={props.color}
      onClick={props.onClick}
    >
      <path d="M9.417 29.083q2.458-1.75 5.062-2.729 2.604-.979 5.521-.979t5.542.979q2.625.979 5.083 2.729Q32.333 27 33.104 24.75q.771-2.25.771-4.75 0-5.875-4-9.875t-9.875-4q-5.875 0-9.875 4t-4 9.875q0 2.5.792 4.75.791 2.25 2.5 4.333ZM20 21.375q-2.417 0-4.083-1.667-1.667-1.666-1.667-4.083 0-2.417 1.667-4.083Q17.583 9.875 20 9.875q2.417 0 4.083 1.667 1.667 1.666 1.667 4.083 0 2.458-1.667 4.104-1.666 1.646-4.083 1.646Zm0 15.292q-3.417 0-6.458-1.313-3.042-1.312-5.313-3.583t-3.583-5.292Q3.333 23.458 3.333 20t1.313-6.479Q5.958 10.5 8.229 8.229t5.292-3.583Q16.542 3.333 20 3.333t6.479 1.313q3.021 1.312 5.292 3.583t3.583 5.292q1.313 3.021 1.313 6.479 0 3.417-1.313 6.458-1.312 3.042-3.583 5.313t-5.292 3.583Q23.458 36.667 20 36.667Zm0-2.792q2.25 0 4.375-.646t4.083-2.187q-1.958-1.375-4.083-2.125T20 28.167q-2.25 0-4.375.75t-4.083 2.125q1.958 1.541 4.083 2.187 2.125.646 4.375.646Zm0-15.25q1.292 0 2.125-.833.833-.834.833-2.167 0-1.292-.833-2.125T20 12.667q-1.292 0-2.125.833t-.833 2.125q0 1.333.833 2.167.833.833 2.125.833Zm0-3Zm0 15.417Z" />
    </svg>
  )
}

export default Add
