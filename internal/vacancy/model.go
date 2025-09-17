package vacancy

import (
	"regexp"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type VacancyFormCreate struct {
	Role     string
	Company  string
	Type     string
	Salary   string
	Location string
	Email    string
}
type Vacancy struct {
	ID        int       `db:"id"`
	Role      string    `db:"role"`
	Company   string    `db:"company"`
	Type      string    `db:"type"`
	Salary    string    `db:"salary"`
	Location  string    `db:"location"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
}

func FromCtx(c *fiber.Ctx) VacancyFormCreate {
	form := VacancyFormCreate{
		Role:     c.FormValue("role"),
		Company:  c.FormValue("company"),
		Type:     c.FormValue("type"),
		Salary:   c.FormValue("salary"),
		Location: c.FormValue("location"),
		Email:    c.FormValue("email"),
	}

	form.sanitize()

	return form
}
func (v *VacancyFormCreate) Validate() map[string]string {
	errors := make(map[string]string)

	// Валидация должности (только буквы, цифры, пробелы и основные символы)
	if strings.TrimSpace(v.Role) == "" {
		errors["role"] = "Должность обязательна для заполнения"
	} else if matched, _ := regexp.MatchString(`^[a-zA-Zа-яА-ЯёЁ0-9\s\-\.\,\(\)]{2,100}$`, v.Role); !matched {
		errors["role"] = "Должность должна содержать от 2 до 100 символов (только буквы, цифры и основные символы)"
	}

	// Валидация названия компании
	if strings.TrimSpace(v.Company) == "" {
		errors["company"] = "Название компании обязательно для заполнения"
	} else if matched, _ := regexp.MatchString(`^[a-zA-Zа-яА-ЯёЁ0-9\s\-\.\,\(\)&]{2,150}$`, v.Company); !matched {
		errors["company"] = "Название компании должно содержать от 2 до 150 символов"
	}

	// Валидация сферы компании
	if strings.TrimSpace(v.Type) == "" {
		errors["type"] = "Сфера компании обязательна для заполнения"
	} else if matched, _ := regexp.MatchString(`^[a-zA-Zа-яА-ЯёЁ\s\-]{2,100}$`, v.Type); !matched {
		errors["type"] = "Сфера компании должна содержать от 2 до 100 символов (только буквы и дефисы)"
	}

	// Валидация заработной платы (формат: цифры, возможно с указанием валюты)
	if strings.TrimSpace(v.Salary) == "" {
		errors["salary"] = "Заработная плата обязательна для заполнения"
	} else if matched, _ := regexp.MatchString(`^[0-9\s\-–—₽€\$₴₸₺₼₾₿₶₷₸₺\.,]{3,20}$`, v.Salary); !matched {
		errors["salary"] = "Укажите корректный формат заработной платы"
	}

	// Валидация расположения
	if strings.TrimSpace(v.Location) == "" {
		errors["location"] = "Расположение обязательно для заполнения"
	} else if matched, _ := regexp.MatchString(`^[a-zA-Zа-яА-ЯёЁ\s\-\.\,]{2,100}$`, v.Location); !matched {
		errors["location"] = "Расположение должно содержать от 2 до 100 символов (только буквы и основные символы)"
	}

	// Валидация email
	if strings.TrimSpace(v.Email) == "" {
		errors["email"] = "Email обязателен для заполнения"
	} else if matched, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, v.Email); !matched {
		errors["email"] = "Укажите корректный email адрес"
	}

	return errors
}

// IsValid проверяет, прошла ли форма валидацию
func (v *VacancyFormCreate) IsValid() bool {
	return len(v.Validate()) == 0
}

// sanitize очищает поля от лишних пробелов
func (v *VacancyFormCreate) sanitize() {
	v.Role = strings.TrimSpace(v.Role)
	v.Company = strings.TrimSpace(v.Company)
	v.Type = strings.TrimSpace(v.Type)
	v.Salary = strings.TrimSpace(v.Salary)
	v.Location = strings.TrimSpace(v.Location)
	v.Email = strings.TrimSpace(strings.ToLower(v.Email))
}

func FormatErrors(errors map[string]string) string {
	var res string
	for _, v := range errors {
		res = res + v + ", "
	}
	return strings.TrimSuffix(res, ", ")
}
