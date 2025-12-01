package sqls

import (
	"testing"
)

func TestCombiners(t *testing.T) {
	tests := []struct {
		name       string
		combiner   Combiner
		conditions []string
		expected   string
	}{
		{
			name:       "And with multiple conditions",
			combiner:   And(),
			conditions: []string{"a = 1", "b = 2", "c = 3"},
			expected:   "a = 1 AND b = 2 AND c = 3",
		},
		{
			name:       "And with single condition",
			combiner:   And(),
			conditions: []string{"a = 1"},
			expected:   "a = 1",
		},
		{
			name:       "And with no conditions",
			combiner:   And(),
			conditions: []string{},
			expected:   "",
		},
		{
			name:       "Or with multiple conditions",
			combiner:   Or(),
			conditions: []string{"a = 1", "b = 2", "c = 3"},
			expected:   "a = 1 OR b = 2 OR c = 3",
		},
		{
			name:       "Or with single condition",
			combiner:   Or(),
			conditions: []string{"a = 1"},
			expected:   "a = 1",
		},
		{
			name:       "Or with no conditions",
			combiner:   Or(),
			conditions: []string{},
			expected:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.combiner.Combine(tt.conditions)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestWhere(t *testing.T) {
	tests := []struct {
		name       string
		conditions []string
		op         Combiner
		expected   string
	}{
		{
			name:       "Where with conditions",
			conditions: []string{"a = 1", "b = 2"},
			op:         And(),
			expected:   "WHERE a = 1 AND b = 2",
		},
		{
			name:       "Where with no conditions",
			conditions: []string{},
			op:         And(),
			expected:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Where(tt.conditions, tt.op)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestHaving(t *testing.T) {
	tests := []struct {
		name       string
		conditions []string
		op         Combiner
		expected   string
	}{
		{
			name:       "Having with conditions",
			conditions: []string{"count(a) > 1", "sum(b) < 100"},
			op:         Or(),
			expected:   "HAVING count(a) > 1 OR sum(b) < 100",
		},
		{
			name:       "Having with no conditions",
			conditions: []string{},
			op:         Or(),
			expected:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Having(tt.conditions, tt.op)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestComparisonOperators(t *testing.T) {
	tests := []struct {
		name     string
		function func([]string, string, string) []string
		field    string
		value    string
		expected string
	}{
		{
			name:     "Eq",
			function: Eq,
			field:    "name",
			value:    "'john'",
			expected: "(name = 'john')",
		},
		{
			name:     "Ne",
			function: Ne,
			field:    "age",
			value:    "18",
			expected: "(age <> 18)",
		},
		{
			name:     "Gt",
			function: Gt,
			field:    "salary",
			value:    "50000",
			expected: "(salary > 50000)",
		},
		{
			name:     "Lt",
			function: Lt,
			field:    "price",
			value:    "100",
			expected: "(price < 100)",
		},
		{
			name:     "Ge",
			function: Ge,
			field:    "score",
			value:    "90",
			expected: "(score >= 90)",
		},
		{
			name:     "Le",
			function: Le,
			field:    "quantity",
			value:    "10",
			expected: "(quantity <= 10)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conditions := tt.function([]string{}, tt.field, tt.value)
			if len(conditions) != 1 {
				t.Fatalf("expected 1 condition, got %d", len(conditions))
			}
			if conditions[0] != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, conditions[0])
			}
		})
	}
}

func TestMustComparisonOperators(t *testing.T) {
	tests := []struct {
		name     string
		function func([]string, string, string) []string
		initial  []string
		field    string
		value    string
		expected []string
	}{
		{
			name:     "MustEq with non-empty value",
			function: MustEq,
			initial:  []string{},
			field:    "name",
			value:    "'john'",
			expected: []string{"(name = 'john')"},
		},
		{
			name:     "MustEq with empty value",
			function: MustEq,
			initial:  []string{"existing_condition"},
			field:    "name",
			value:    "",
			expected: []string{"existing_condition"},
		},
		{
			name:     "MustNe with non-empty value",
			function: MustNe,
			initial:  []string{},
			field:    "age",
			value:    "18",
			expected: []string{"(age <> 18)"},
		},
		{
			name:     "MustNe with empty value",
			function: MustNe,
			initial:  []string{"existing_condition"},
			field:    "age",
			value:    "",
			expected: []string{"existing_condition"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conditions := tt.function(tt.initial, tt.field, tt.value)
			if len(conditions) != len(tt.expected) {
				t.Fatalf("expected %d conditions, got %d", len(tt.expected), len(conditions))
			}
			for i, expected := range tt.expected {
				if conditions[i] != expected {
					t.Errorf("condition %d: expected %q, got %q", i, expected, conditions[i])
				}
			}
		})
	}
}

func TestRangeOperators(t *testing.T) {
	tests := []struct {
		name     string
		function func([]string, string, string, string) []string
		field    string
		minValue string
		maxValue string
		expected string
		mustTest bool
	}{
		{
			name:     "Between",
			function: func(conds []string, field, min, max string) []string { return Between(conds, field, min, max) },
			field:    "age",
			minValue: "18",
			maxValue: "65",
			expected: "(age BETWEEN 18 AND 65)",
		},
		{
			name:     "NotBetween",
			function: func(conds []string, field, min, max string) []string { return NotBetween(conds, field, min, max) },
			field:    "score",
			minValue: "0",
			maxValue: "59",
			expected: "(score NOT BETWEEN 0 AND 59)",
		},
		{
			name:     "MustBetween with valid values",
			function: func(conds []string, field, min, max string) []string { return MustBetween(conds, field, min, max) },
			field:    "age",
			minValue: "18",
			maxValue: "65",
			expected: "(age BETWEEN 18 AND 65)",
			mustTest: true,
		},
		{
			name:     "MustBetween with empty min value",
			function: func(conds []string, field, min, max string) []string { return MustBetween(conds, field, min, max) },
			field:    "age",
			minValue: "",
			maxValue: "65",
			expected: "",
			mustTest: true,
		},
		{
			name:     "MustNotBetween with valid values",
			function: func(conds []string, field, min, max string) []string { return MustNotBetween(conds, field, min, max) },
			field:    "score",
			minValue: "0",
			maxValue: "59",
			expected: "(score NOT BETWEEN 0 AND 59)",
			mustTest: true,
		},
		{
			name:     "MustNotBetween with empty max value",
			function: func(conds []string, field, min, max string) []string { return MustNotBetween(conds, field, min, max) },
			field:    "score",
			minValue: "0",
			maxValue: "",
			expected: "",
			mustTest: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var conditions []string
			if tt.mustTest {
				conditions = tt.function([]string{}, tt.field, tt.minValue, tt.maxValue)
			} else {
				conditions = tt.function([]string{}, tt.field, tt.minValue, tt.maxValue)
			}

			if tt.expected == "" {
				if len(conditions) != 0 {
					t.Errorf("expected no conditions, got %v", conditions)
				}
			} else {
				if len(conditions) != 1 {
					t.Fatalf("expected 1 condition, got %d", len(conditions))
				}
				if conditions[0] != tt.expected {
					t.Errorf("expected %q, got %q", tt.expected, conditions[0])
				}
			}
		})
	}
}

func TestLikeOperators(t *testing.T) {
	tests := []struct {
		name     string
		function func([]string, string, string) []string
		field    string
		value    string
		expected string
		mustTest bool
	}{
		{
			name:     "Like",
			function: Like,
			field:    "name",
			value:    "%john%",
			expected: "(name LIKE '%john%')",
		},
		{
			name:     "NotLike",
			function: NotLike,
			field:    "email",
			value:    "%.spam.com",
			expected: "(email NOT LIKE '%.spam.com')",
		},
		{
			name:     "MustLike with non-empty value",
			function: MustLike,
			field:    "description",
			value:    "%urgent%",
			expected: "(description LIKE '%urgent%')",
			mustTest: true,
		},
		{
			name:     "MustLike with empty value",
			function: MustLike,
			field:    "description",
			value:    "",
			expected: "",
			mustTest: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var conditions []string
			if tt.mustTest {
				conditions = tt.function([]string{}, tt.field, tt.value)
			} else {
				conditions = tt.function([]string{}, tt.field, tt.value)
			}

			if tt.expected == "" {
				if len(conditions) != 0 {
					t.Errorf("expected no conditions, got %v", conditions)
				}
			} else {
				if len(conditions) != 1 {
					t.Fatalf("expected 1 condition, got %d", len(conditions))
				}
				if conditions[0] != tt.expected {
					t.Errorf("expected %q, got %q", tt.expected, conditions[0])
				}
			}
		})
	}
}

func TestInOperators(t *testing.T) {
	tests := []struct {
		name     string
		function func([]string, string, []string) []string
		field    string
		values   []string
		expected string
		mustTest bool
	}{
		{
			name:     "In",
			function: In,
			field:    "id",
			values:   []string{"1", "2", "3"},
			expected: "(id IN (1,2,3))",
		},
		{
			name:     "NotIn",
			function: NotIn,
			field:    "status",
			values:   []string{"'inactive'", "'deleted'"},
			expected: "(status NOT IN ('inactive','deleted'))",
		},
		{
			name:     "MustIn with non-empty values",
			function: MustIn,
			field:    "category",
			values:   []string{"'books'", "'electronics'"},
			expected: "(category IN ('books','electronics'))",
			mustTest: true,
		},
		{
			name:     "MustIn with empty values",
			function: MustIn,
			field:    "category",
			values:   []string{},
			expected: "",
			mustTest: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var conditions []string
			if tt.mustTest {
				conditions = tt.function([]string{}, tt.field, tt.values)
			} else {
				conditions = tt.function([]string{}, tt.field, tt.values)
			}

			if tt.expected == "" {
				if len(conditions) != 0 {
					t.Errorf("expected no conditions, got %v", conditions)
				}
			} else {
				if len(conditions) != 1 {
					t.Fatalf("expected 1 condition, got %d", len(conditions))
				}
				if conditions[0] != tt.expected {
					t.Errorf("expected %q, got %q", tt.expected, conditions[0])
				}
			}
		})
	}
}

func TestNullOperators(t *testing.T) {
	tests := []struct {
		name     string
		function func([]string, string) []string
		field    string
		expected string
	}{
		{
			name:     "IsNull",
			function: IsNull,
			field:    "updated_at",
			expected: "(updated_at IS NULL)",
		},
		{
			name:     "IsNotNull",
			function: IsNotNull,
			field:    "created_at",
			expected: "(created_at IS NOT NULL)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conditions := tt.function([]string{}, tt.field)
			if len(conditions) != 1 {
				t.Fatalf("expected 1 condition, got %d", len(conditions))
			}
			if conditions[0] != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, conditions[0])
			}
		})
	}
}

func TestExistentialOperators(t *testing.T) {
	tests := []struct {
		name     string
		function func([]string, string) []string
		subquery string
		expected string
	}{
		{
			name:     "Exists",
			function: Exists,
			subquery: "SELECT 1 FROM orders o WHERE o.customer_id = c.id",
			expected: "(EXISTS (SELECT 1 FROM orders o WHERE o.customer_id = c.id))",
		},
		{
			name:     "NotExists",
			function: NotExists,
			subquery: "SELECT 1 FROM banned_users b WHERE b.user_id = u.id",
			expected: "(NOT EXISTS (SELECT 1 FROM banned_users b WHERE b.user_id = u.id))",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conditions := tt.function([]string{}, tt.subquery)
			if len(conditions) != 1 {
				t.Fatalf("expected 1 condition, got %d", len(conditions))
			}
			if conditions[0] != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, conditions[0])
			}
		})
	}
}

func TestSubqueryComparisonOperators(t *testing.T) {
	tests := []struct {
		name     string
		function func([]string, string, string) []string
		field    string
		subquery string
		expected string
	}{
		{
			name:     "EqAll",
			function: EqAll,
			field:    "department_id",
			subquery: "SELECT id FROM departments WHERE active = 1",
			expected: "(department_id = ALL (SELECT id FROM departments WHERE active = 1))",
		},
		{
			name:     "EqAny",
			function: EqAny,
			field:    "salary",
			subquery: "SELECT avg_salary FROM roles r WHERE r.role = e.role",
			expected: "(salary = ANY (SELECT avg_salary FROM roles r WHERE r.role = e.role))",
		},
		{
			name:     "NeAll",
			function: NeAll,
			field:    "status",
			subquery: "SELECT status FROM blacklisted_statuses",
			expected: "(status <> ALL (SELECT status FROM blacklisted_statuses))",
		},
		{
			name:     "NeAny",
			function: NeAny,
			field:    "category_id",
			subquery: "SELECT id FROM excluded_categories",
			expected: "(category_id <> ANY (SELECT id FROM excluded_categories))",
		},
		{
			name:     "GtAll",
			function: GtAll,
			field:    "score",
			subquery: "SELECT min_score FROM grade_thresholds WHERE grade = 'A'",
			expected: "(score > ALL (SELECT min_score FROM grade_thresholds WHERE grade = 'A'))",
		},
		{
			name:     "GtAny",
			function: GtAny,
			field:    "price",
			subquery: "SELECT avg_price FROM product_categories pc WHERE pc.id = p.category_id",
			expected: "(price > ANY (SELECT avg_price FROM product_categories pc WHERE pc.id = p.category_id))",
		},
		{
			name:     "LtAll",
			function: LtAll,
			field:    "age",
			subquery: "SELECT max_age FROM age_groups WHERE group = 'minor'",
			expected: "(age < ALL (SELECT max_age FROM age_groups WHERE group = 'minor'))",
		},
		{
			name:     "LtAny",
			function: LtAny,
			field:    "experience",
			subquery: "SELECT min_exp FROM job_levels jl WHERE jl.level = e.level",
			expected: "(experience < ANY (SELECT min_exp FROM job_levels jl WHERE jl.level = e.level))",
		},
		{
			name:     "GeAll",
			function: GeAll,
			field:    "rating",
			subquery: "SELECT threshold FROM quality_benchmarks qb WHERE qb.type = p.type",
			expected: "(rating >= ALL (SELECT threshold FROM quality_benchmarks qb WHERE qb.type = p.type))",
		},
		{
			name:     "GeAny",
			function: GeAny,
			field:    "discount",
			subquery: "SELECT standard_discount FROM customer_tiers ct WHERE ct.tier = c.tier",
			expected: "(discount >= ANY (SELECT standard_discount FROM customer_tiers ct WHERE ct.tier = c.tier))",
		},
		{
			name:     "LeAll",
			function: LeAll,
			field:    "quantity",
			subquery: "SELECT stock_limit FROM warehouses w WHERE w.id = i.warehouse_id",
			expected: "(quantity <= ALL (SELECT stock_limit FROM warehouses w WHERE w.id = i.warehouse_id))",
		},
		{
			name:     "LeAny",
			function: LeAny,
			field:    "priority",
			subquery: "SELECT default_priority FROM user_preferences up WHERE up.user_id = t.assignee_id",
			expected: "(priority <= ANY (SELECT default_priority FROM user_preferences up WHERE up.user_id = t.assignee_id))",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conditions := tt.function([]string{}, tt.field, tt.subquery)
			if len(conditions) != 1 {
				t.Fatalf("expected 1 condition, got %d", len(conditions))
			}
			if conditions[0] != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, conditions[0])
			}
		})
	}
}

func TestChainingConditions(t *testing.T) {
	conditions := []string{}

	// Chain several conditions together
	conditions = Eq(conditions, "name", "'john'")
	conditions = Gt(conditions, "age", "18")
	conditions = Like(conditions, "email", "%@example.com")

	if len(conditions) != 3 {
		t.Fatalf("expected 3 conditions, got %d", len(conditions))
	}

	expected := []string{
		"(name = 'john')",
		"(age > 18)",
		"(email LIKE '%@example.com')",
	}

	for i, exp := range expected {
		if conditions[i] != exp {
			t.Errorf("condition %d: expected %q, got %q", i, exp, conditions[i])
		}
	}

	// Test combining them with And()
	result := And().Combine(conditions)
	expectedResult := "(name = 'john') AND (age > 18) AND (email LIKE '%@example.com')"
	if result != expectedResult {
		t.Errorf("expected %q, got %q", expectedResult, result)
	}
}
