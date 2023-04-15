package schema

type Role string

const (
	User      Role = "user"
	Assistant Role = "assistant"
	System    Role = "system"
)

type Message struct {
	Name    string `json:"name"`
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

func UserMessage(content string) Message {
	return Message{Role: User, Content: content}
}

func AssistantMessage(content string) Message {
	return Message{Role: Assistant, Content: content}
}

func SystemMessage(content string) Message {
	return Message{Role: System, Content: content}
}

func ChatMessage(role string, content string) Message {
	return Message{Role: Role(role), Content: content}
}
