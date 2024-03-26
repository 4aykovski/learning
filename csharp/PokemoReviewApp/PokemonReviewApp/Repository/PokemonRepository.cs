using PokemonReviewApp.Interface;
using PokemoReviewApp.Data;
using PokemoReviewApp.Models;

namespace PokemonReviewApp.Repository;

public class PokemonRepository : IPokemonRepository
{
    private readonly DataContext _context;

    public PokemonRepository(DataContext context)
    {
        _context = context;
    }

    public ICollection<Pokemon> GetPokemons()
    {
        return _context.Pokemons.OrderBy(p => p.Id).ToList();
    }
}