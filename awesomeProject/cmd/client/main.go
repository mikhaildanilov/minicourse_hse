package main

import (
	"awesomeProject/accounts/dto"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

type Command struct {
	Port      int
	Host      string
	Cmd       string
	Name      string
	NameNew   string
	Amount    int
	AmountNew int
}

func main() {
	portVal := flag.Int("port", 8080, "server port")
	hostVal := flag.String("host", "0.0.0.0", "server host")
	cmdVal := flag.String("cmd", "", "command to execute")
	nameVal := flag.String("name", "", "name of account")
	amountVal := flag.Int("amount", 0, "amount of account")
	nameNewVal := flag.String("nameNew", "", "new name of account")
	amountNewVal := flag.Int("amountNew", 0, "updated amount of account")

	flag.Parse()

	cmd := Command{
		Port:      *portVal,
		Host:      *hostVal,
		Cmd:       *cmdVal,
		Name:      *nameVal,
		Amount:    *amountVal,
		NameNew:   *nameNewVal,
		AmountNew: *amountNewVal,
	}

	if err := cmd.do(); err != nil {
		panic(err)
	}
}

func (c *Command) do() error {
	switch c.Cmd {
	case "create":
		return c.create()
	case "get":
		return c.get()
	case "delete":
		return c.deleteAccount()
	case "changeName":
		return c.changeName()
	case "changeAmount":
		return c.changeBalance()
	default:
		return fmt.Errorf("unknown command: %s", c.Cmd)
	}
}

func (c *Command) create() error {
	request := dto.CreateAccountRequest{
		Name:   c.Name,
		Amount: c.Amount,
	}

	data, err := json.Marshal(request)

	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/create", c.Host, c.Port),
		"application/json",
		bytes.NewReader(data),
	)

	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusCreated {
		return nil
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func (c *Command) get() error {
	resp, err := http.Get(
		fmt.Sprintf("http://%s:%d/account?name=%s", c.Host, c.Port, c.Name),
	)

	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("read body failed: %w", err)
		}

		return fmt.Errorf("resp error %s", string(body))
	}

	var response dto.GetAccountResponse

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("json decode failed: %w", err)
	}

	fmt.Printf("response account name: %s ; amount: %d", response.Name, response.Amount)

	return nil
}

func (c *Command) deleteAccount() error {
	request := dto.DeleteAccountRequest{
		Name: c.Name,
	}

	data, err := json.Marshal(request)

	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/deleteAccount", c.Host, c.Port),
		"application/json",
		bytes.NewReader(data),
	)

	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func (c *Command) changeName() error {
	request := dto.ChangeAccountNameRequest{
		Name:    c.Name,
		NameNew: c.NameNew,
	}

	data, err := json.Marshal(request)

	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/changeName", c.Host, c.Port),
		"application/json",
		bytes.NewReader(data),
	)

	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func (c *Command) changeBalance() error {
	request := dto.ChangeAccountBalanceRequest{
		Name:      c.Name,
		AmountNew: c.AmountNew,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/changeBalance", c.Host, c.Port),
		"application/json",
		bytes.NewReader(data),
	)

	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}
