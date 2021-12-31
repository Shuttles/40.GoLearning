package gopher_test

import (
	"github.com/Shuttles/40.Golearning/GinkgoExample/gopher"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func mockInputData() ([]gopher.Gopher, error) {
	inputData := []gopher.Gopher{
		{
			Name:   "菜刀",
			Gender: "男",
			Age:    18,
		},
		{
			Name:   "小西瓜",
			Gender: "女",
			Age:    19,
		},
		{
			Name:   "机器铃砍菜刀",
			Gender: "男",
			Age:    17,
		},
		{
			Name:   "小菜刀",
			Gender: "男",
			Age:    20,
		},
	}
	return inputData, nil
}

var _ = Describe("Gopher", func() {
	var (
		inputData []gopher.Gopher
		err       error
	)
	BeforeEach(func() {
		inputData, err = mockInputData()
		By("当测试不通过时，我会在这里打印一个消息 [BeforeEach]")
	})

	AfterEach(func() {
		By("当测试不通过时，我会在这里打印一个信息 [AfterEach]")
	})

	Describe("校验输入数据", func() {
		Context("当获取数据没有错误发生时", func() {
			It("应该是接收数据成功了的", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("当获取的数据校验失败时", func() {
			When("当数据校验检查到gopher名字太短，小于3时", func() {
				It("应该返回错误：名字太短，不能小于3", func() {
					Expect(gopher.Validate(inputData[0])).To(MatchError("名字太短，不能小于3"))
				})
			})
			When("当数据校验检查到gopher性别为女时", func() {
				It("应该返回错误：抱歉！不是男生，目前只面向男生", func() {
					Expect(gopher.Validate(inputData[1])).To(MatchError("抱歉！不是男生，目前只面向男生"))
				})
			})
			When("当数据校验检查到gopher年龄太小时", func() {
				It("应该返回错误：年龄太小，不能小于18", func() {
					Expect(gopher.Validate(inputData[2])).To(MatchError("年龄太小，不能小于18"))
				})
			})
		})

		Context("当获取的数据校验成功时", func() {
			It("通过了数据校验", func() {
				Expect(gopher.Validate(inputData[3])).To(BeNil())
			})
		})
	})

})
