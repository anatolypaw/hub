export default function Button({ children }) {
    function handleClick() {
        console.log("clicked" + children)
    }
    return (
        <button onClick={handleClick}>{children}</button>
    )
}