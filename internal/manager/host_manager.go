package manager

import (
    "fmt"
    "regexp"
    "strings"

    "github.com/will-wright-eng/sshm/internal/config"
    "github.com/will-wright-eng/sshm/internal/models"
)

type HostManager struct {
    config config.SSHConfigReader
}

func NewHostManager(config config.SSHConfigReader) *HostManager {
    return &HostManager{config: config}
}

func (m *HostManager) formatHostEntry(entry *models.HostEntry) string {
    return fmt.Sprintf(`Host %s
  Hostname %s
  User %s
  IdentityFile %s
  Port %d
`, entry.Name, entry.Hostname, entry.User, entry.IdentityFile, entry.Port)
}

func (m *HostManager) AddHost(entry *models.HostEntry) error {
    content, err := m.config.Read()
    if err != nil {
        return fmt.Errorf("reading config: %w", err)
    }

    if !strings.Contains(content, config.DelimiterStart) {
        content = fmt.Sprintf("%s\n\n%s\n%s\n", content, config.DelimiterStart, config.DelimiterEnd)
    }

    newHost := m.formatHostEntry(entry)
    re := regexp.MustCompile(fmt.Sprintf(`(%s\n)(.*?)(\n%s)`, config.DelimiterStart, config.DelimiterEnd))
    updatedContent := re.ReplaceAllString(content, "${1}${2}"+newHost+"${3}")

    return m.config.Write(updatedContent)
}

func (m *HostManager) RemoveHost(name string) error {
    content, err := m.config.Read()
    if err != nil {
        return fmt.Errorf("reading config: %w", err)
    }

    // Create regex to match the host entry
    hostPattern := fmt.Sprintf(`Host %s\n(?:  [^\n]+\n)*`, regexp.QuoteMeta(name))
    re := regexp.MustCompile(hostPattern)

    // Check if host exists in managed block
    managedBlock := m.getManagedBlock(content)
    if !re.MatchString(managedBlock) {
        return fmt.Errorf("host %s not found in managed block", name)
    }

    // Remove the host entry
    updatedBlock := re.ReplaceAllString(managedBlock, "")
    updatedContent := m.replaceManagedBlock(content, updatedBlock)

    return m.config.Write(updatedContent)
}

func (m *HostManager) ListHosts() ([]*models.HostEntry, error) {
    content, err := m.config.Read()
    if err != nil {
        return nil, fmt.Errorf("reading config: %w", err)
    }

    managedBlock := m.getManagedBlock(content)
    if managedBlock == "" {
        return nil, nil
    }

    var hosts []*models.HostEntry
    hostRe := regexp.MustCompile(`Host ([^\n]+)\n  Hostname ([^\n]+)\n  User ([^\n]+)\n  IdentityFile ([^\n]+)\n  Port (\d+)`)

    matches := hostRe.FindAllStringSubmatch(managedBlock, -1)
    for _, match := range matches {
        port := 22
        fmt.Sscanf(match[5], "%d", &port)

        host := models.NewHostEntry(
            match[1], // name
            match[2], // hostname
            match[3], // user
            match[4], // identity file
            port,
        )
        hosts = append(hosts, host)
    }

    return hosts, nil
}

// Helper methods
func (m *HostManager) getManagedBlock(content string) string {
    re := regexp.MustCompile(fmt.Sprintf(`%s\n(.*?)\n%s`, config.DelimiterStart, config.DelimiterEnd))
    match := re.FindStringSubmatch(content)
    if len(match) < 2 {
        return ""
    }
    return match[1]
}

func (m *HostManager) replaceManagedBlock(content, newBlock string) string {
    re := regexp.MustCompile(fmt.Sprintf(`(%s\n).*?(\n%s)`, config.DelimiterStart, config.DelimiterEnd))
    return re.ReplaceAllString(content, "${1}"+newBlock+"${2}")
}
