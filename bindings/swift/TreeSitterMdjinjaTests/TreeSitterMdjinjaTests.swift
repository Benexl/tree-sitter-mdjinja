import XCTest
import SwiftTreeSitter
import TreeSitterMdjinja

final class TreeSitterMdjinjaTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_mdjinja())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading mdjinja grammar")
    }
}
