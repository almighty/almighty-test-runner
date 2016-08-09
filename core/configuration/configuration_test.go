package configuration_test

import (
    cfg "github.com/almighty/almighty-test-runner/core/configuration"
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)


var _ = Describe("Test runner configuration", func() {

    Context("Loading build settings", func() {

        It("Should load build command", func() {
            // given
            configurationAsYAML := `
            build_tool:
                cmd: mvn clean install -P all
            `

            // when
            configuration, _ := cfg.Read(configurationAsYAML)

            // then
            Expect(configuration.Command).To(Equal("mvn clean install -P all"))
        })

        It("Should load description", func() {
            // given
            configurationAsYAML := `
            description: "I'm GROOT!"
            build_tool:
                cmd: mvn clean install -P all
            `

            // when
            configuration, _ := cfg.Read(configurationAsYAML)

            // then
            Expect(configuration.Description).To(Equal("I'm GROOT!"))
        })

        It("Should load path", func() {
            // given
            configurationAsYAML := `
            path: /tmp/project/
            `

            // when
            configuration, _ := cfg.Read(configurationAsYAML)

            // then
            Expect(configuration.Path).To(Equal("/tmp/project/"))
        })

    })

    Context("Enviroment verticals", func() {
        It("Should load verticals", func() {
            // given
            configurationAsYAML := `
            build_tool:
                cmd: mvn clean test
            verticals:
                database: [hsqldb, oracle11g, mysql, postgresql, mssql]
                runtime: [dropwizard, wildfly-swarm, express.js, goa]
            `

            // when
            configuration, _ := cfg.Read(configurationAsYAML)

            // when
            Expect(configuration.Verticals).To(HaveKeyWithValue("database", []string {"hsqldb", "oracle11g", "mysql", "postgresql", "mssql"}))
            Expect(configuration.Verticals).To(HaveKeyWithValue("runtime", []string {"dropwizard", "wildfly-swarm", "express.js", "goa"}))
        })

        It("Should provide cartesian product of all verticals", func() {
            // given
            configurationAsYAML := `
            verticals:
                database: [hsqldb, oracle11g, mysql, postgresql, mssql]
                runtime: [dropwizard, wildfly-swarm, express.js, goa]
            `
            configuration, _ := cfg.Read(configurationAsYAML)

            // when
            products := cfg.ProductFor(cfg.VerticalSelection{Vertical: "runtime", Selection: "dropwizard"}, configuration.EnvironmentVerticals)

            // when
            Expect(products).To(HaveLen(5))
            Expect(products).To(ConsistOf([][]string{
                {"hsqldb", "dropwizard"},
                {"oracle11g", "dropwizard"},
                {"mysql", "dropwizard"},
                {"postgresql", "dropwizard"},
                {"mssql", "dropwizard"},
            }))
        })

        It("Should provide combination of other verticles with selected verticle", func() {
            // given
            configurationAsYAML := `
            verticals:
                database: [hsqldb, oracle11g, mysql, postgresql, mssql]
                runtime: [dropwizard, wildfly-swarm, express.js, goa]
            `
            configuration, _ := cfg.Read(configurationAsYAML)

            // when
            products := cfg.Product(configuration.EnvironmentVerticals)

            // when
            Expect(products).To(HaveLen(len(allProducts())))
            Expect(products).To(ConsistOf(allProducts()))
        })
    })

})

// -- Test fixtures
func allProducts() [][]string {
    return [][]string{
        {"hsqldb", "dropwizard"},
        {"hsqldb", "wildfly-swarm"},
        {"hsqldb", "express.js"},
        {"hsqldb", "goa"},
        {"oracle11g", "dropwizard"},
        {"oracle11g", "wildfly-swarm"},
        {"oracle11g", "express.js"},
        {"oracle11g", "goa"},
        {"mysql", "dropwizard"},
        {"mysql", "wildfly-swarm"},
        {"mysql", "express.js"},
        {"mysql", "goa"},
        {"postgresql", "dropwizard"},
        {"postgresql", "wildfly-swarm"},
        {"postgresql", "express.js"},
        {"postgresql", "goa"},
        {"mssql", "dropwizard"},
        {"mssql", "wildfly-swarm"},
        {"mssql", "express.js"},
        {"mssql", "goa"},

    }
}
